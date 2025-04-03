package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/Raipus/ZoomerOK/account/pkg/grpc/pb"
	"github.com/Raipus/ZoomerOK/account/pkg/postgres"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAccountServer
}

func (s *server) Authorization(ctx context.Context, req *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	s := strings.Split(req.Token, ":")
	if s[0] != "Token" || !postgres.UUIDExists(s[1]) {
		return &pb.AuthorizationResponse{
			Success: false,
		}, nil
	}

	return &pb.AuthorizationResponse{
		Success: true,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthorizationServer(s, &server{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
