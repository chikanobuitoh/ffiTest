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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

//export get
func get(mes *C.char) *C.char {
	tex := C.GoString(mes)

	hostaddr, err := readJsonFile("Sample.ini", "host")
	//errの時はファイルが無い時なのでlocalhostにつなぎます
	if err == nil {
		fmt.Println("connect to :" + hostaddr)
		addr = &hostaddr
	} else {
		fmt.Println("missing is Sample.ini :+ " + err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleSerciveClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := "gRPC Check"
	r, err := c.Check(ctx, &pb.CheckRequest{Request: req})
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

func readJsonFile(path string, readname string) (string, error) {

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	var Data interface{}
	if err := json.Unmarshal([]byte(raw), &Data); err != nil {
		return "", err
	}

	outPut := Data.(map[string]interface{})[readname].(string)
	return outPut, nil
}
