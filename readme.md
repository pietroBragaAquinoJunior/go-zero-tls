goctl rpc protoc ./protos/zrpc-um.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc-um

goctl rpc protoc ./protos/zrpc-dois.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc-dois


# Chave e Certificado Auto-assinado
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./tls/server.key -out ./tls/server.crt -config ./tls/openssl.cnf


# Fazer a chamada pro método com TLS
grpcurl -cacert tls/server.crt -import-path ./protos -proto zrpc-dois.proto -d "{\"Nome\":\"pietro\"}" localhost:8081 usuario.ZrpcDoisService.ZrpcDoisMethod

grpcurl -cacert tls/server.crt -import-path ./protos -proto zrpc-um.proto -d "{\"Nome\":\"pietro\"}" localhost:8080 usuario.ZrpcUmService.ZrpcUmMethod