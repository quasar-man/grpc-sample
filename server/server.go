package server

import (
	"context"

	"github.com/quasar-man/grpc-sample/pb"
)

type SampleServer struct {
	pb.UnimplementedGreetingServiceServer
}

func NewSampleServer() *SampleServer {
	return &SampleServer{}
}


func (s *SampleServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}
