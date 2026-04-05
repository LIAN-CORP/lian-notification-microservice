package bootstrap

import (

	"go.uber.org/zap"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lian-corp/notification-microservice/internal/application"
	"github.com/lian-corp/notification-microservice/internal/config"
	"github.com/lian-corp/notification-microservice/internal/domain/api"
	"github.com/lian-corp/notification-microservice/internal/domain/api/usecase"
	"github.com/lian-corp/notification-microservice/internal/domain/spi"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/adapter"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/repository"
	kafkaconsumer "github.com/lian-corp/notification-microservice/internal/infrastructure/driving/kafka/consumer"
)

type Container struct {
	NotificationConsumer *kafkaconsumer.KafkaConsumer
	Log					 *zap.SugaredLogger
	Topics 				 []string
}

func NewContainer(cfg *config.Config) *Container {
	container, err := buildContainer(cfg)
	if err != nil {
		panic(err)
	}
	return container
}

func buildContainer(cfg *config.Config) (*Container, error) {
	
	// ========== Logger ==========
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sugar := log.Sugar()

	// ========== Database ==========
	db, err := postgres.NewConnection(cfg.DB.URL, sugar)
	if err != nil {
		return nil, err
	}

	// ========== Repositories ==========
	notifcationEventRepo := *repository.NewEventRepository(db)

	// ========== Adapters ==========
	//var notificationEventAdapter *adapter.NotificationEventAdapter = adapter.NewNotificationEventAdapter(notifcationEventRepo)
	var notificationEventPersistence spi.NotificationEventPersistence = adapter.NewNotificationEventAdapter(notifcationEventRepo)
	

	// ========== UseCases ==========
	var notificationEventService api.NotificationEventService = usecase.NewNotificationEventUseCase(notificationEventPersistence)
	var notificationService api.NotificationService = usecase.NewNotificationServiceUseCase(notificationEventService)

	// ========== Handlers ==========
	handler := application.NewNotificationConsumerHandler(notificationService)

	// ========== Kafka Consumer ==========
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.BootstrapServers,
		"group.id":			 cfg.Kafka.GroupID,
		"auto.offset.reset": cfg.Kafka.AutoOffsetReset,
	})

	if err != nil {
		return nil, err
	}

	consumer := kafkaconsumer.NewKafkaConsumer(kafkaConsumer, handler, sugar, cfg.Kafka.Topics)

	// ========== Container ==========
	return &Container{
		Log: sugar,
		NotificationConsumer: consumer,
	}, nil

}