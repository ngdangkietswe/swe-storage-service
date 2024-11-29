package grpc

import (
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/grpc/middleware"
	"github.com/ngdangkietswe/swe-storage-service/configs"
	"github.com/ngdangkietswe/swe-storage-service/grpc/server/storage"
	"github.com/ngdangkietswe/swe-storage-service/minio"
	"google.golang.org/grpc"
	"log"
	"net"
)

// NewGrpcServer function is used to create a new gRPC server. It listens on the gRPC port and serves the gRPC server.
func NewGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.GlobalConfig.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	minioClient := minio.NewMinIO()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthMiddleware),
	)

	storage.NewGrpcHandler(minioClient).RegisterGrpcServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
