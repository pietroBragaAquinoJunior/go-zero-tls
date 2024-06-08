package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"

	"go-zero-tls/common/pb"
	"go-zero-tls/zrpc-dois/internal/config"
	"go-zero-tls/zrpc-dois/internal/server"
	"go-zero-tls/zrpc-dois/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/zrpcdois.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	creds, err := credentials.NewServerTLSFromFile("../server.crt", "../server.key")
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterZrpcDoisServiceServer(grpcServer, server.NewZrpcDoisServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()
	s.AddOptions(grpc.Creds(creds))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
