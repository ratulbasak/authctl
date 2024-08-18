package service

import "encoding/json"

type HelloService struct{}

type HelloMessage struct {
	Message string `json:"message"`
}

func (s *HelloService) GetHello() ([]byte, error) {
	message := HelloMessage{
		Message: "hello world",
	}
	return json.Marshal(message)
}
