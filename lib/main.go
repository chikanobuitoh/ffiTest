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
		log.Fatalf("could not greet: %v", err)
	}

	return C.CString(tex + "_" + r.GetResult())
}

func main() {
}
