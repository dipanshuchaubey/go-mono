package utils

import (
	"carthage/services/gateway/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type HttpResponseStruct struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func PopulateErrorRespose(err error) []byte {
	var httpRes HttpResponseStruct

	httpRes.Error = err.Error()

	jsonData, err := json.Marshal(httpRes)

	if err != nil {
		fmt.Println("Error converting data to JSON: ", err.Error())
		return nil
	}

	return jsonData
}

func PopulateSuccessRespose(data interface{}) []byte {
	httpRes := HttpResponseStruct{
		Data:    data,
		Success: true,
	}

	jsonData, err := json.Marshal(httpRes)

	if err != nil {
		fmt.Println("Error converting data to JSON: ", err.Error())
		return nil
	}

	return jsonData
}

func SetParamsInContext(ctx context.Context, cnf types.Route, r *http.Request) context.Context {
	if strings.Contains(cnf.URL, "{") {
		paramArr := strings.Split(cnf.URL, "/")
		rawParam := paramArr[len(paramArr)-1]
		param := strings.Trim(rawParam, "{}")

		ctx = context.WithValue(ctx, param, r.PathValue(param))
	}

	return ctx
}

func Middleware(req http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Application-Type", "application/json")
				w.Write(PopulateErrorRespose(fmt.Errorf("%v", err)))
			}
		}()

		req.ServeHTTP(w, r)
	})
}
