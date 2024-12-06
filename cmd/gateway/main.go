package main

import (
	common "carthage/common/config"
	"carthage/services/gateway/constants"
	"carthage/services/gateway/routes"
	"carthage/services/gateway/types"
	"carthage/services/gateway/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	ServiceName = "gateway"
)

func main() {
	// auth.Auth()

	// Load env based config
	var env types.Config
	common.LoadConfig(ServiceName, &env)

	configs := routes.RouteConfig(&env)
	h := routes.NewRouteMap(&env)

	for url, route := range configs {
		ctx := context.Background()
		url := constants.API_URL_PREFIX + url

		handler := func(w http.ResponseWriter, r *http.Request) {
			for _, cnf := range route {
				ctx := utils.SetParamsInContext(ctx, cnf, r)

				handlerCaller, found := h[cnf.Handler]
				if !found {
					fmt.Println("Handler not found: ", cnf.Handler)
					continue
				}

				if cnf.Method == r.Method {
					data, err := handlerCaller(ctx, r)

					w.Header().Set("Content-Type", "application/json")

					if err != nil {
						w.Write(utils.PopulateErrorRespose(err))
					} else {
						w.Write(utils.PopulateSuccessRespose(data))
					}
				}
			}

		}

		http.Handle(url, utils.Middleware(http.HandlerFunc(handler)))

		for _, cnf := range route {
			fmt.Printf("Registered route for %s: %s %s\n", cnf.Handler, cnf.Method, url)
		}
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", env.Service.Port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
