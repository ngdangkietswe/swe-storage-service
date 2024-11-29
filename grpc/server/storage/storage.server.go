package storage

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
	storageservice "github.com/ngdangkietswe/swe-storage-service/grpc/service/storage"
)

type GrpcServer struct {
	storage.UnimplementedStorageServiceServer
	storageService storageservice.IStorageService
}

func NewGrpcServer(storageService storageservice.IStorageService) *GrpcServer {
	return &GrpcServer{
		storageService: storageService,
	}
}

// GetPresignedURL is a function that implements the GetPresignedURL method of the StorageServiceServer interface
func (s *GrpcServer) GetPresignedURL(ctx context.Context, req *storage.PresignedURLReq) (*storage.PresignedURLResp, error) {
	return s.storageService.GetPresignedURL(ctx, req)
}
