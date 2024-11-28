package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

// PutObj uploads an object to the specified bucket.
func (c *Client) PutObj(ctx context.Context, bucket, object, filePath string) {
	data, err := c.minioClient.FPutObject(ctx, bucket, object, filePath, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	if err != nil {
		log.Printf("[MINIO] Failed to put object: %v", err)
		return
	}

	log.Printf("[MINIO] Successfully uploaded %s of size %d to %s", object, data.Size, bucket)
}

// PresignedUrl generates a presigned URL for the specified object.
func (c *Client) PresignedUrl(ctx context.Context, bucket, object string) string {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+object+"\"")
	presignedUrl, err := c.minioClient.PresignedGetObject(ctx, bucket, object, time.Second*time.Duration(configs.GlobalConfig.MinIODefaultExpiry), reqParams)
	if err != nil {
		log.Printf("[MINIO] Failed to get presigned URL: %v", err)
		return ""
	}

	log.Printf("[MINIO] Successfully generated presigned URL: %s", presignedUrl.String())
	return presignedUrl.String()
}
