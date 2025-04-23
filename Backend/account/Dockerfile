FROM golang:1.24
WORKDIR /home/app/account

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o account cmd/main.go