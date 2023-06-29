# blockChain-Service

生成表 ddl下
goctl model mysql ddl --src mall.sql --dir ./myssql
生成 客户端服务端 block包下
goctl rpc protoc block.proto --go_out=. --go-grpc_out=. --zrpc_out=.