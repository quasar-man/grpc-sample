package server

import (
	"context"
	"errors"
	"fmt"
	"io"
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

func (s *SampleServer) HelloClientStream(stream pb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			// クライアントからのリクエストが終了した時の処理
			messages :=  fmt.Sprintf("Hello, %v", nameList)
			return stream.SendAndClose(&pb.HelloResponse{Message: messages})
		}

		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}
