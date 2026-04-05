package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Kafka KafkaConfig
	DB DBConfig
	AWS AWSConfig
	App AppConfig
}

type KafkaConfig struct {
	BootstrapServers string
	GroupID string
	AutoOffsetReset string
	Topics []string
}

type DBConfig struct {
	URL string
}

type AWSConfig struct {
	Region string
}

type AppConfig struct {
	AppPort string
}

func Load() *Config {

	//Load .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//DB Url
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "lian_notification")

	connectionURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	kafkaTopics := []string{"notification.debt.events"}

	return &Config{
		Kafka: KafkaConfig{
			BootstrapServers: getEnv("KAFKA_BROKERS", "localhost:9092"),
			GroupID: getEnv("KAFKA_TOPIC", "debt.event"),
			AutoOffsetReset: getEnv("KAFKA_AUTO_OFFSET_RESET", "earliest"),
			Topics: kafkaTopics,
		},
		DB: DBConfig{
			URL: connectionURL,
		},
		AWS: AWSConfig{
			Region: getEnv("AWS_REGION", "us-east-1"),
		},
		App: AppConfig{
			AppPort: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}