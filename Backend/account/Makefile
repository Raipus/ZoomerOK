start:
	go run cmd/main.go

test:
	go test -v -cover github.com/Raipus/ZoomerOK/account/pkg/handlers

proto:
	protoc -I ./pkg/grpc/pb/ --go_out=. --go-grpc_out=. ./pkg/grpc/pb/*.proto