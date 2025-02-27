package server

import (
	"context"
	"github.com/ngdangkietswe/swe-go-common-shared/constants"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-go-common-shared/security"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
	storageservice "github.com/ngdangkietswe/swe-storage-service/grpc/service/storage"
)

type StorageGrpcServer struct {
	storage.UnimplementedStorageServiceServer
	storageService storageservice.IStorageService
}

func NewStorageGrpcServer(storageService storageservice.IStorageService) *StorageGrpcServer {
	return &StorageGrpcServer{
		storageService: storageService,
	}
}

// GetPresignedURL is a function that implements the GetPresignedURL method of the StorageServiceServer interface
func (s *StorageGrpcServer) GetPresignedURL(ctx context.Context, req *storage.PresignedURLReq) (*storage.PresignedURLResp, error) {
	var action string
	if req.Type == storage.PresignedURLMethod_PRESIGNED_URL_METHOD_GET {
		action = constants.ActionDownload
	} else {
		action = constants.ActionUpload
	}
	return security.SecuredAuth(ctx, req, domain.Permission{Action: action, Resource: constants.ResourceFile}, s.storageService.GetPresignedURL)
}
