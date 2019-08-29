package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bookun/sandbox/go/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	host := ""
	conn, err := grpc.Dial(host+":19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	//client := pb.NewCatClient(conn)
	//message := &pb.GetMyCatMessage{TargetCat: "tama"}
	//res, err := client.GetMyCat(context.TODO(), message)
	client := pb.NewHelloServiceClient(conn)
	message := &pb.HelloRequest{Name: "hoge"}
	res, err := client.Hello(context.TODO(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result:%v \n", res)
}
