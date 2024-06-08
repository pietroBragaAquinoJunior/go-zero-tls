package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	__ "go-zero-tls/common/pb"
	"go-zero-tls/zrpc-um/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	ZrpcDoisClient __.ZrpcDoisServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	zrpcDoisClient := zrpc.MustNewClient(c.ZrpcDoisClientConf)
	return &ServiceContext{
		Config:         c,
		ZrpcDoisClient: __.NewZrpcDoisServiceClient(zrpcDoisClient.Conn()),
	}
}
