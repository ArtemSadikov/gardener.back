protoc:
	protoc -I=./proto/user -I=. \
	--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./internal/api/grpc/interface \
	--go_out=paths=source_relative:./internal/api/grpc/interface \
	./proto/user/user.proto

protoc_sdk:
	protoc -I=./proto/user -I=../../pkg/sdk/users \
	--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:../../pkg/sdk/users/interface \
	--go_out=paths=source_relative:../../pkg/sdk/users/interface \
	./proto/user/user.proto

build_api_grpc: protoc
	go build -o build/api/grpc ./cmd/api/main.go

docker:
	docker build -t users-service .
