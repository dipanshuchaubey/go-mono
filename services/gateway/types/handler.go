package types

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, req *http.Request) (interface{}, error)
