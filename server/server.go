package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/afa4/fullcycle-go-grpc/pb/protos"
	"google.golang.org/grpc"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (*GreeterService) SayHello(ctx context.Context, helloRequest *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: fmt.Sprintf("hello %s", helloRequest.Name)}, nil
}

func (*GreeterService) Speech(helloRequest *pb.HelloRequest, serv pb.Greeter_SpeechServer) error {
	for {
		time.Sleep(time.Second * 1)
		err := serv.Send(&pb.HelloReply{Message: "Talking..."})
		if err != nil {
			break
		}

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic("Failed to start socket")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &GreeterService{})
	err = grpcServer.Serve(lis)
	if err != nil {
		panic("Failed to start server")
	}
}
