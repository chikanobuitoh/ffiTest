package main

import (
	"C"
	"ffimodule/pb"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

//export ffiCheck
func ffiCheck(mes *C.char) *C.char {
	tex := C.GoString(mes)

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleSerciveClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Check(ctx, &pb.CheckRequest{Request: "gRPC Check"})
	if err != nil {
		log.Fatalf("Check Fail: %v", err)
	}

	//Streamで受け取ってファイル化
	basePath := "sample.png"
	file, err := os.Create(basePath)
	if err != nil {
		fmt.Println(err)
		return C.CString("cant open file")
	}
	defer file.Close()

	var (
		feedbackDAT *pb.FeedBack
	)

	replystring := ""

	for {
		reply, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return C.CString("could not stream")
		}

		if reply.Result != "" {
			replystring = reply.Result
		}

		if reply.Feedback != nil {
			feedbackDAT = reply.GetFeedback()
			file.Write(feedbackDAT.ResponseFile)
		}
	}

	file.Close()

	return C.CString(tex + "_" + replystring)
}

func main() {
}
