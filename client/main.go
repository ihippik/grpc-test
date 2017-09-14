package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	pb "github.com/ihippik/grpc-test/protocol"
	"github.com/golang/protobuf/proto"
)

func main(){
	res, err :=http.Get("http://localhost:8070")
	if err != nil{
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err.Error())
	}

	var u pb.User
	err = proto.Unmarshal(b, &u)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(u)
}
