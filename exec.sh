goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc/ --go-grpc_out=./apps/user/rpc/ --zrpc_out=./apps/user/rpc/

goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/users/models/" -c
goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero

goctl api go -api apps/meeting/api/meeting.api -dir apps/meeting/api -style gozero
goctl rpc protoc ./apps/meeting/rpc/meeting.proto --go_out=./apps/meeting/rpc/ --go-grpc_out=./apps/meeting/rpc/ --zrpc_out=./apps/meeting/rpc/

goctl api go -api apps/ws/ws.api -dir apps/ws/api -style gozero