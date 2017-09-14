package main

import (
	"net/http"
	pb "github.com/ihippik/grpc-test/protocol"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", send)
	e.Logger.Fatal(e.Start(":1323"))
}

func send(c echo.Context) error {
	u:= &pb.User{
		Name: "Константин",
		Email: "hippik80@gmail.com",
		Id:1,
	}

	body, err :=proto.Marshal(u)
	if err!=nil{
		return c.String(http.StatusInternalServerError,err.Error())
	}
	return c.Blob(http.StatusOK, "application/x-protobuf", body)
}