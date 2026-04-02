package main

import (
	"context"
	"fmt"

	"github.com/lian-corp/notification-microservice/internal/config"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driving/kafka/consumer"
)

type Server struct {
	consumer *consumer.KafkaConsumer
	msgCH chan string
}

func NewServer() *Server {
	msgCH := make(chan string, 64)
	c, err := consumer.NewKafkaConsumer(msgCH)
	
	if err != nil {
		panic("Failed to create Kafka consumer: " + err.Error())
	}

	return &Server{
		consumer: c,
		msgCH: msgCH,
	}
}

func (s *Server) handleMsg(msg string) {
	//db operation
	fmt.Printf("received msg: %s\n", msg)
}

func main() {
	app := config.New()
	logger := config.NewLogger()
	logger.Log.Sync()

	err := app.Start(context.TODO())
	if err != nil {
		logger.Sugar.Error("Failed to start the app: ", err)
	}

	s := NewServer()
	for msg := range s.msgCH {
		go s.handleMsg(msg)
	}
}

