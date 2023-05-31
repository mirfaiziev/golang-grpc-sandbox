package grpc

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/mirfaiziev/golang-grpc-sandbox/internal/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHelloWorld(ctx context.Context, in *emptypb.Empty) (*pb.HelloWorldReply, error) {
	return &pb.HelloWorldReply{Message: "Hello word - 1"}, nil
}

func StartGRPCServer() {
	var grpcServerPort = os.Getenv("GRPC_SERVER_PORT")
	listener, err := net.Listen("tcp", ":"+grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Start listening port %s", grpcServerPort)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Printf("Start grpc server")
	}
}
