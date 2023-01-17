package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"ffimodule/pb"

	"google.golang.org/grpc"
	refrection "google.golang.org/grpc/reflection"
)

// Port設定
var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSampleSerciveServer
}

func (s *server) Check(in *pb.CheckRequest, stream pb.SampleSercive_CheckServer) error {

	fileData := "./server/sample.png"
	//Responceを返す
	file, err := os.Open(fileData)
	if err != nil {
		fmt.Println("Cant FileOpen :" + err.Error())
		return err
	}
	defer file.Close()

	stream.Send(&pb.CheckResponce{
		Result: in.GetRequest(),
	})

	for {
		convertFileToBinary, err := convertFileToBinary(file)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("failed upload file: " + err.Error())
			return err
		}

		stream.Send(&pb.CheckResponce{
			Feedback: &pb.FeedBack{
				ResponseFile: convertFileToBinary,
			},
		})
	}

	file.Close()

	return nil
}

func convertFileToBinary(file *os.File) (fileBinary []byte, err error) {
	uploadFileBinary := make([]byte, 1024)
	count, err := file.Read(uploadFileBinary)

	log.Printf("file successfully loaded", count)

	if err != nil {
		return nil, err
	}

	return uploadFileBinary, nil
}

func main() {
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

	pb.RegisterSampleSerciveServer(s, &server{})
	refrection.Register(s)
	log.Printf("server listening at %v", listenPort.Addr())
	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	s.Serve(listenPort)
}
