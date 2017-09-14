package main

import (
	"net/http"
	pb "github.com/ihippik/grpc-test/protocol"

	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request){
	u:= &pb.User{
		Id: proto.Int32(1234),
		Name: proto.String("Константин"),
		Email: proto.String("hippik80@gmail.com"),
	}

	body, err :=proto.Marshal(u)
	if err!=nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)
}