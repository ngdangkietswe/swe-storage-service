package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/storage"
	"github.com/ngdangkietswe/swe-storage-service/configs"
	"log"
	"net/url"
	"time"
)

type Client struct {
	minioClient *minio.Client
}

func NewMinIO() *Client {
	minioClient, err := minio.New(
		fmt.Sprintf("%s:%d", configs.GlobalConfig.MinIOHost, configs.GlobalConfig.MinIOPort),
		&minio.Options{
			Creds:  credentials.NewStaticV4(configs.GlobalConfig.MinIOAccessKey, configs.GlobalConfig.MinIOSecretKey, ""),
			Secure: configs.GlobalConfig.MinIOUseSSL,
		})

	if err != nil {
		log.Printf("[MINIO] Failed to create minio client: %v", err)
		return nil
	}

	return &Client{
		minioClient: minioClient,
	}
}

// PresignedUrl generates a presigned URL for the specified object.
// The presigned URL can be used to perform PUT or GET operations on the object.
func (c *Client) PresignedUrl(ctx context.Context, bucket, objectName string, method storage.PresignedURLMethod, duration int32) (string, error) {
	if method != storage.PresignedURLMethod_PRESIGNED_URL_METHOD_PUT && method != storage.PresignedURLMethod_PRESIGNED_URL_METHOD_GET {
		return "", fmt.Errorf("[MINIO] Invalid presigned URL method: %v", method)
	}

	expires := time.Second * time.Duration(configs.GlobalConfig.MinIODefaultExpiry)
	if duration > 0 {
		expires = time.Second * time.Duration(duration)
	}

	var (
		presignedUrl *url.URL
		err          error
	)

	switch method {
	case storage.PresignedURLMethod_PRESIGNED_URL_METHOD_PUT:
		presignedUrl, err = c.minioClient.PresignedPutObject(ctx, bucket, objectName, expires)
	case storage.PresignedURLMethod_PRESIGNED_URL_METHOD_GET:
		reqParams := url.Values{
			"response-content-disposition": {fmt.Sprintf("attachment; filename=\"%s\"", objectName)},
		}
		presignedUrl, err = c.minioClient.PresignedGetObject(ctx, bucket, objectName, expires, reqParams)
	}

	if err != nil {
		return "", fmt.Errorf("[MINIO] failed to generate presigned URL: %w", err)
	}

	log.Printf("[MINIO] Successfully generated presigned URL: %s", presignedUrl.String())
	return presignedUrl.String(), nil
}
