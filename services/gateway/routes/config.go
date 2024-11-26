package routes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	Route struct {
		URL       string `yaml:"url"`
		Method    string `yaml:"method"`
		Handler   string `yaml:"handler"`
		Timeout   int    `yaml:"timeout"`
		Auth      bool   `yaml:"auth"`
		RateLimit int    `yaml:"rateLimit"`
	}

	Routes map[string]Route

	HandlerFunc func(ctx context.Context, req *http.Request) (interface{}, error)
)

func readRoutesYAML(path string) (*Routes, error) {
	var routes *Routes

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

func ReadConfig() *Routes {
	conf, routeErr := readRoutesYAML("/Users/dipanshu/Developer/private/go-mono/services/gateway/routes/routes.yaml")

	if routeErr != nil {
		fmt.Printf("Error reading routes: %s\n", routeErr)
		os.Exit(1)
	}

	return conf
}
