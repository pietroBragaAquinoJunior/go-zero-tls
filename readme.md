goctl rpc protoc ./protos/zrpc-um.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc-um

goctl rpc protoc ./protos/zrpc-dois.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc-dois

# Gerar chave privada
openssl genrsa -out server.key 2048

# Gerar certificado autoassinado
openssl req -new -x509 -key server.key -out server.crt -days 365
