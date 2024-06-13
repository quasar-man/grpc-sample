package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/quasar-man/grpc-sample/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/quasar-man/grpc-sample/server"
)

func main() {
	port := 8080

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPC Server 作成
	s := grpc.NewServer()

	// gRPC Server にサービス(GreetingService)を登録
	pb.RegisterGreetingServiceServer(s, server.NewSampleServer())

	// サーバーリフレクションの設定
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}