package main

import (
	"context"
	"github.com/ngdangkietswe/swe-storage-service/configs"
	"github.com/ngdangkietswe/swe-storage-service/minio"
)

func main() {
	minioClient := minio.NewMinIO()
	ctx := context.Background()
	// Test: get presigned URL for the object "test/aws.png"
	minioClient.PresignedUrl(ctx, configs.GlobalConfig.MinIOBucketPublic, "test/aws.png")
}
