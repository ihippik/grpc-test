package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/ihippik/grpc-test/protocol"
)

func main() {
	u := &user.User{
		Name:   "Konstantin",
		Email: "hippik80@gmail.com",
		Id:1,
	}
	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatal("marshaling error: ", err)
		return
	}

	u2 := new(user.User)
	err = proto.Unmarshal(data, u2)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Println(u2)
}