package main

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-storage-service/grpc"
)

func main() {
	config.Init()
	grpc.NewGrpcServer()
}
