package storage

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
	"github.com/ngdangkietswe/swe-storage-service/minio"
)

type storageService struct {
	minioClient *minio.Client
}

// GetPresignedURL returns a presigned URL for the object in the bucket.
func (s storageService) GetPresignedURL(ctx context.Context, req *storage.PresignedURLReq) (*storage.PresignedURLResp, error) {
	url, err := s.minioClient.PresignedUrl(ctx, req.BucketName, req.ObjectName, req.Type, req.Duration)
	if err != nil {
		return nil, err
	}

	return &storage.PresignedURLResp{
		Success: true,
		Resp: &storage.PresignedURLResp_Url{
			Url: url,
		},
	}, nil
}

func NewStorageService(minioClient *minio.Client) IStorageService {
	return &storageService{
		minioClient: minioClient,
	}
}
