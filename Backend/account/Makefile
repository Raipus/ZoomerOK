start:
	go run cmd/main.go

proto:
	protoc -I ./pkg/grpc/pb/ --go_out=. --go-grpc_out=. ./pkg/grpc/pb/*.proto