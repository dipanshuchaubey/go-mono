package main

import (
	ot "carthage/common/otel"
	us "carthage/protos/user_service"
	service "carthage/services/user_service"
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

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	name   = "user-service"
	tracer = otel.Tracer(name)
	meter  = otel.Meter(name)
	logger = otelslog.NewLogger(name)
)

func timeoutInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// Create a channel to catch the result or a panic recovery
	done := make(chan struct{})
	var resp interface{}
	var err error

	// Set up OpenTelemetry.
	otelShutdown, otelErr := ot.SetupOTelSDK(ctx)
	if otelErr != nil {
		return nil, otelErr
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic in %s: %v", info.FullMethod, r)
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
	var wg sync.WaitGroup
	wg.Add(1)

	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered from panic:", err)
			}
		}()

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen on port 50051: %v", err)
		}

		options := []grpc.ServerOption{
			grpc.ConnectionTimeout(time.Second),
			grpc.UnaryInterceptor(timeoutInterceptor),
		}

		s := grpc.NewServer(options...)

		us.RegisterUserServiceServer(s, service.NewUserService())

		reflection.Register(s)

		go func() {
			log.Printf("gRPC server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()

		<-stop // Wait for stop signal
		log.Println("Shutting down gRPC server...")
		s.GracefulStop()
	}()

	wg.Wait()
}
