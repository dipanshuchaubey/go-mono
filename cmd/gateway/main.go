package main

import (
	"carthage/services/gateway/config"
	"carthage/services/gateway/constants"
	"carthage/services/gateway/routes"
	"carthage/services/gateway/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// auth.Auth()

	configs := routes.ReadConfig()
	h := config.NewConf()

	for _, cnf := range *configs {
		handlerCaller, found := h[cnf.Handler]
		if !found {
			fmt.Println("Handler not found: ", cnf.Handler)
			continue
		}

		ctx := context.Background()
		url := constants.API_URL_PREFIX + cnf.URL

		handler := func(w http.ResponseWriter, r *http.Request) {
			ctx := utils.SetParamsInContext(ctx, cnf, r)

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

		http.Handle(url, utils.Middleware(http.HandlerFunc(handler)))

		fmt.Printf("Registered route for %s: %s %s\n", cnf.Handler, cnf.Method, url)
	}

	err := http.ListenAndServe(":5000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
