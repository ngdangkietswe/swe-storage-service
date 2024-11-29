package storage

import (
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
	service "github.com/ngdangkietswe/swe-storage-service/grpc/service/storage"
	"github.com/ngdangkietswe/swe-storage-service/minio"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
	minioClient *minio.Client
}

func NewGrpcHandler(client *minio.Client) *GrpcHandler {
	return &GrpcHandler{
		minioClient: client,
	}
}

func (h *GrpcHandler) RegisterGrpcServer(server *grpc.Server) {
	storageService := service.NewStorageService(h.minioClient)
	storage.RegisterStorageServiceServer(server, NewGrpcServer(storageService))
}
