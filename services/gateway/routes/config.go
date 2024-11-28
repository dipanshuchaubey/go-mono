package routes

import (
	"carthage/services/gateway/constants"
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

func ReadConfig() *types.Routes {
	conf, routeErr := readRoutesYAML(constants.ROUTES_CONFIG_FILE_PATH)

	if routeErr != nil {
		fmt.Printf("Error reading routes: %s\n", routeErr)
		os.Exit(1)
	}

	return conf
}

func NewConf(config *types.Config) map[string]types.HandlerFunc {
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
