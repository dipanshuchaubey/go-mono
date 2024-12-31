package main

import (
	common "carthage/common/config"
	ot "carthage/common/otel"
	service "carthage/services/bootcamp_service"
	"carthage/services/bootcamp_service/config"
	"carthage/services/bootcamp_service/constants"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	bs "github.com/dipanshuchaubey/protos-package/bootcamp_service"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var Tracer = otel.Tracer(constants.ServiceName)
var logger = otelslog.NewLogger(constants.ServiceName)

func timeoutInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// Create a channel to catch the result or a panic recovery
	done := make(chan struct{})
	var resp interface{}
	var err error

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.InfoContext(ctx, fmt.Sprintf("Recovered from panic in %s: %v", info.FullMethod, r))
				err = status.Error(codes.Internal, "internal server error")
			}
			close(done)
		}()
		resp, err = handler(ctx, req)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		// Handle context timeout
		log.Printf("Timeout in %s", info.FullMethod)
		return nil, status.Error(codes.DeadlineExceeded, "request timed out")
	case <-done:
		// Proceed normally if no timeout
		return resp, err
	}
}

func main() {
	otelShutdown, err := ot.SetupOTelSDK()
	if err != nil {
		return
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err := errors.Join(err, otelShutdown(context.Background()))
		if err != nil {
			logger.Error(fmt.Sprintf("Error shutting down OTEL: %v", err))
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Load config
	var config config.Config
	common.LoadConfig(constants.ServicePathName, &config)

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				logger.Warn(fmt.Sprintf("recovered from panic: %v", err))
			}
		}()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Service.Port))
		if err != nil {
			logger.Error(fmt.Sprintf("failed to listen on port %s: %v", config.Service.Port, err))
		}

		options := []grpc.ServerOption{
			grpc.ConnectionTimeout(time.Second),
			grpc.UnaryInterceptor(timeoutInterceptor),
		}

		s := grpc.NewServer(options...)

		bs.RegisterBootcampServiceServer(s, service.NewBootcampService(config))

		reflection.Register(s)

		go func() {
			logger.Info(fmt.Sprintf("gRPC server listening at %v", lis.Addr()))
			if err := s.Serve(lis); err != nil {
				logger.Error(fmt.Sprintf("failed to serve: %v", err))
			}
		}()

		<-stop // Wait for stop signal
		logger.Info("Shutting down gRPC server...")
		s.GracefulStop()
	}()

	wg.Wait()
}
