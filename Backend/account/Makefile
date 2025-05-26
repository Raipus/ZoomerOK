start:
	go run cmd/main.go

test:
	go test -v -cover github.com/Raipus/ZoomerOK/account/pkg/handlers

verify:
	go mod verify

proto:
	protoc -I ./pkg/broker/pb/ --go_out=. ./pkg/broker/pb/*.proto

lint:
	golangci-lint run