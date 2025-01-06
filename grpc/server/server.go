package server

import (
	"context"
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/middleware"
	grpcstorage "github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Params struct {
	fx.In
	StorageGrpcServer *StorageGrpcServer
}

// NewGrpcServer function is used to create a new gRPC server. It listens on the gRPC port and serves the gRPC server.
func NewGrpcServer(lifecycle fx.Lifecycle, params Params) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthMiddleware),
	)
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetInt("GRPC_PORT", 7030)))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			grpcstorage.RegisterStorageServiceServer(grpcServer, params.StorageGrpcServer)

			go func() {
				log.Printf("gRPC server is running on %s", lis.Addr().String())
				if err = grpcServer.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("gRPC server is shutting down...")
			grpcServer.GracefulStop()
			return nil
		},
	})
	return grpcServer
}
