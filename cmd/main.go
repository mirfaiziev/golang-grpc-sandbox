package main

import (
	"context"
	pb "github.com/mirfaiziev/golang-grpc-sandbox/internal/pb"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHelloWorld(ctx context.Context, in *emptypb.Empty) (*pb.HelloWorldReply, error) {
	return &pb.HelloWorldReply{Message: "Hello word"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
