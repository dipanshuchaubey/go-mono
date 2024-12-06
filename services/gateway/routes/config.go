package routes

import (
	"carthage/services/gateway/handlers"
	"carthage/services/gateway/types"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func readRoutesYAML(path string) (*types.Routes, error) {
	var routes *types.Routes

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	errUnmar := yaml.Unmarshal(fileBytes, &routes)
	if errUnmar != nil {
		return nil, errUnmar
	}

	return routes, nil
}

func RouteConfig(config *types.Config) types.RoutesMap {
	conf, routeErr := readRoutesYAML(config.Service.RoutesConfigPath)

	if routeErr != nil {
		fmt.Printf("Error reading routes: %s\n", routeErr)
		os.Exit(1)
	}

	var routesMap = make(types.RoutesMap)

	for _, cnf := range *conf {
		if _, ok := routesMap[cnf.URL]; !ok {
			routesMap[cnf.URL] = make(map[string]types.Route)
			routesMap[cnf.URL][cnf.Method] = cnf
		} else {
			routesMap[cnf.URL][cnf.Method] = cnf
		}
	}

	return routesMap
}

func NewRouteMap(config *types.Config) map[string]types.HandlerFunc {
	bootcampSvc := handlers.BootcampHandler(config)
	userSvc := handlers.UserHandler(config)

	HandlersMap := map[string]types.HandlerFunc{
		// Bootcamp Service
		"GetBootcamps":   bootcampSvc.GetBootcamps(),
		"CreateBootcamp": bootcampSvc.CreateBootcamp(),
		// User Service
		"GetUsers": userSvc.GetUsers(),
		"GetUser":  userSvc.GetUser(),
	}

	return HandlersMap
}
