package storage

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
)

type IStorageService interface {
	GetPresignedURL(ctx context.Context, req *storage.PresignedURLReq) (*storage.PresignedURLResp, error)
}
