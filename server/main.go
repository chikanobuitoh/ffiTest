package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"pb"

	"google.golang.org/grpc"
	refrection "google.golang.org/grpc/reflection"
)

// Port設定
var (
	port = flag.Int("port", 80, "The server port")
)

type server struct {
	pb.UnimplementedSampleSerciveServer
}

func (s *server) Check(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponce, error) {
	return &pb.CheckResponce{
		Result: "notExist",
	}, nil
}

func Service() {
	flag.Parse()
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		dir = "Null"
	}
	fmt.Println(dir)

	s := grpc.NewServer()

	pb.RegisterGiftPleaseSerciveServer(s, &server{})
	refrection.Register(s)
	log.Printf("server listening at %v", listenPort.Addr())
	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	s.Serve(listenPort)
}
