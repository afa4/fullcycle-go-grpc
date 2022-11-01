package main

import (
	"context"
	"fmt"

	pb "github.com/afa4/fullcycle-go-grpc/pb/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("Failed to dial")
	}
	defer conn.Close()
	grpcClient := pb.NewGreeterClient(conn)
	reply, err := grpcClient.SayHello(context.Background(), &pb.HelloRequest{Name: "Arthur"})
	fmt.Println(reply)
	stream, err := grpcClient.Speech(context.Background(), &pb.HelloRequest{Name: "Arthur"})
	for {
		reply, err = stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(reply)
	}
}
