package main

import (
	"C"
	"flag"
	"log"

	"pb"

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
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Check(ctx, &pb.CheckRequest{Result: tex})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return C.CString(r)
}

func main() {
}

type BookDatabase struct{}
type Book struct {
	Bookname string "hogehoge"
}

func NewBookDatabase() *BookDatabase {
	var db BookDatabase
	return &db
}

func (d *BookDatabase) GetBook(id string) *Book {
	return &Book{Bookname: "okok"}
}

func (d *BookDatabase) CreateBook(data interface{}) {}

type BookReader interface {
	GetBook(string) *Book
}

func fetch(b BookReader, id string) *Book {
	return b.GetBook(id)
}
