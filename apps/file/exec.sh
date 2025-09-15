goctl rpc protoc ./apps/file/rpc/file.proto --go_out=./apps/file/rpc/ --go-grpc_out=./apps/file/rpc/ --zrpc_out=./apps/file/rpc/

goctl api go -api apps/file/api/file.api -dir apps/file/api -style gozero