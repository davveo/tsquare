package main

import (
	"context"
	"fmt"

	uuidproto "github.com/davveo/tsquare/proto/uuid"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService()

	service.Init()

	cl := uuidproto.NewUuidService("go.micro.srv.uuid", service.Client())

	resp, err := cl.GenerateId(context.Background(), &uuidproto.Request{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
