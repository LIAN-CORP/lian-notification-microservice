package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lian-corp/notification-microservice/internal/application"
	"go.uber.org/zap"
)

type KafkaConsumer struct {
	kafkaConsumer 	*kafka.Consumer
	kafkaHandler 	*application.NotificationConsumerHandler
	log 			*zap.SugaredLogger
	topics 			[]string
}

func NewKafkaConsumer(
	consumer *kafka.Consumer, 
	handler *application.NotificationConsumerHandler, 
	log *zap.SugaredLogger,
	topics []string,
) *KafkaConsumer {
	return &KafkaConsumer{
		kafkaConsumer: consumer,
		kafkaHandler: handler,
		log: log,
		topics: topics,
	}
}

func (c *KafkaConsumer) StartConsumer(ctx context.Context) {
	for {
		c.kafkaConsumer.SubscribeTopics(c.topics, nil)
		msg, err := c.kafkaConsumer.ReadMessage(-1)
		if err != nil{
			c.log.Fatalw("Error reading message from kafka", "error", err)
			continue
		}

		if err := c.kafkaHandler.HandleMessage(msg.Value); err != nil {
			c.log.Fatalw("Error handling message", "error", err)
		}
	}
}
