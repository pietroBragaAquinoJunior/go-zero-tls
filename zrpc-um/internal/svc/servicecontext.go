package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	__ "zrpc-aula-2/common/pb"
	"zrpc-aula-2/zrpc-um/internal/config"
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
