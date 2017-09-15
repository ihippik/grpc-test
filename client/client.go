package main

import (
	"fmt"
	pb "github.com/ihippik/grpc-test/protocol"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

func main(){
	address := "localhost:55555"
	conn, err :=grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
	}
	defer conn.Close()
	c:= pb.NewGetUserClient(conn)
	id:=int64(1)
	hr:=&pb.GetUserRequest{Id:id}
	r, err := c.Get(context.Background(), hr)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(r)
}
