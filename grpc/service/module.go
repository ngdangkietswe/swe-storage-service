package service

import (
	grpcservicestorage "github.com/ngdangkietswe/swe-storage-service/grpc/service/storage"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	grpcservicestorage.NewStorageGrpcService,
)
