package config

import "os"

type KafkaConfig struct {
	Topic string
	ConsumerGroup string
	Host string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Topic: os.Getenv("KAFKA_TOPIC"),
		ConsumerGroup: os.Getenv("KAFKA_CONSUMER_GROUP"),
		Host: os.Getenv("KAFKA_HOST"),
	}
}