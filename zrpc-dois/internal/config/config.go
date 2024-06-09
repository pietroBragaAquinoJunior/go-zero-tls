package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	TLS TLSConfig
}

type TLSConfig struct {
	CertFile string
	KeyFile  string
}
