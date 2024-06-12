package server

import (
	"context"
	"fmt"
	"time"

	"github.com/quasar-man/grpc-sample/pb"
)

type SampleServer struct {
	pb.UnimplementedGreetingServiceServer
}

func NewSampleServer() *SampleServer {
	return &SampleServer{}
}


func (s *SampleServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Ohayo, konnichiha, konbanha " + req.GetName()}, nil
}

func (s *SampleServer) HelloServerStream(req *pb.HelloRequest, stream pb.GreetingService_HelloServerStreamServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(
			&pb.HelloResponse{Message: fmt.Sprintf("[%d] Ohayo, konnichiha, konbanha, %v", i,req.GetName())}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}
