package main

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-storage-service/grpc"
	"github.com/ngdangkietswe/swe-storage-service/minio"
	"go.uber.org/fx"
	grpcserver "google.golang.org/grpc"
)

func main() {
	config.Init()
	app := fx.New(
		minio.Module,
		grpc.Module,
		fx.Invoke(func(*grpcserver.Server) {}),
	)
	app.Run()
}
