package main

import (
	"fmt"
	"time"

	"github.com/davveo/tsquare/lib/mq/rabbitmq"
)

func main() {
	var handler = rabbitmq.MQService{}
	go SendSomething(handler)
	handler.Read(tt)
}

func tt(jsonStr []byte) {
	fmt.Println("already in listenning mq: ", string(jsonStr))
}

func SendSomething(handler rabbitmq.MQ) {
	time.Sleep(time.Second * 1)
	handler.Delay("testkey1", "testvalue1", "3000")

	time.Sleep(time.Second * 1)
	handler.Delay("testKey2", "testValue2", "4000")

	time.Sleep(time.Second * 1)
	handler.Delay("testKey3", "testValue3", "5000")

	handler.Delay("delayKey1", "this is delay key233", "6000")
}
