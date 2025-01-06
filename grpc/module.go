package grpc

import (
	"github.com/ngdangkietswe/swe-storage-service/grpc/server"
	"github.com/ngdangkietswe/swe-storage-service/grpc/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	service.Module,
	server.Module,
)
