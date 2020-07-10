package rabbitmq

type MQService struct {
	Rabbitmq
}

type MQ interface {
	Read(func(jsonStr []byte))
	Send(key, value string)
	Delay(key, value, expire string)
}
