package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	__ "go-zero-tls/common/pb"
	"go-zero-tls/zrpc-um/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type ServiceContext struct {
	Config         config.Config
	ZrpcDoisClient __.ZrpcDoisServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	creds, err := credentials.NewClientTLSFromFile(c.TLS.CertFile, "")
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}
	zrpcDoisClient := zrpc.MustNewClient(c.ZrpcDoisClientConf, zrpc.WithDialOption(grpc.WithTransportCredentials(creds)))
	return &ServiceContext{
		Config:         c,
		ZrpcDoisClient: __.NewZrpcDoisServiceClient(zrpcDoisClient.Conn()),
	}
}
