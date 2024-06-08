package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"

	"go-zero-tls/common/pb"
	"go-zero-tls/zrpc-um/internal/config"
	"go-zero-tls/zrpc-um/internal/server"
	"go-zero-tls/zrpc-um/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/zrpcum.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	creds, err := credentials.NewServerTLSFromFile("../cert.pem", "../key.pem")
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterZrpcUmServiceServer(grpcServer, server.NewZrpcUmServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()

	s.AddOptions(grpc.Creds(creds))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	s.Start()
}
