package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	uuid_proto "github.com/zbrechave/tsquare/srv/uuid/proto/uuid"
)

func main() {
	service := micro.NewService()

	service.Init()

	cl := uuid_proto.NewUuidService("go.micro.srv.uuid", service.Client())

	resp, err := cl.GenerateId(context.Background(), &uuid_proto.Request{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
