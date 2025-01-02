package main

import (
	common "carthage/common/config"
	ot "carthage/common/otel"
	"carthage/services/gateway/constants"
	"carthage/services/gateway/routes"
	"carthage/services/gateway/types"
	"carthage/services/gateway/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

const (
	ServiceName = "gateway"
)

var (
	Tracer = otel.Tracer(constants.ServiceName)
	logger = otelslog.NewLogger(constants.ServiceName)
)

func main() {
	// auth.Auth()

	otelShutdown, otErr := ot.SetupOTelSDK()
	if otErr != nil {
		return
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err := errors.Join(otErr, otelShutdown(context.Background()))
		if err != nil {
			logger.Error(fmt.Sprintf("Error shutting down OTEL: %v", err))
		}
	}()

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
					logger.ErrorContext(ctx, fmt.Sprintf("Handler not found: %v", cnf.Handler))
					continue
				}

				if cnf.Method == r.Method {
					ctx, span := Tracer.Start(ctx, "handler."+cnf.Handler)
					defer span.End()

					data, err := handlerCaller(ctx, r)
					setResponse(&w, data, err)
				}
			}

		}

		http.Handle(url, utils.Middleware(http.HandlerFunc(handler)))

		for _, cnf := range route {
			logger.InfoContext(ctx, fmt.Sprintf("Registered route for %s: %s %s", cnf.Handler, cnf.Method, url))
		}
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", env.Service.Port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("server closed")
	} else if err != nil {
		logger.Error(fmt.Sprintf("error starting server: %v", err))
		os.Exit(1)
	}
}

func setResponse(w *http.ResponseWriter, data interface{}, err error) {
	(*w).Header().Set("Content-Type", "application/json")

	var response []byte

	if err != nil {
		response = utils.PopulateErrorResponse(err)
	} else {
		response = utils.PopulateSuccessResponse(data)
	}

	if _, err := (*w).Write(response); err != nil {
		logger.Error(fmt.Sprintf("Error writing response: %v", err))
	}
}
