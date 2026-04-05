package adapter

import (
	"github.com/lian-corp/notification-microservice/internal/domain/model"
	"github.com/lian-corp/notification-microservice/internal/domain/spi"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/mapper"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/repository"
)

type NotificationEventAdapter struct {
	notificationEventRepository repository.NotificationEventRepository
}

func NewNotificationEventAdapter(notificationEventRepository repository.NotificationEventRepository) *NotificationEventAdapter {
	return &NotificationEventAdapter{notificationEventRepository: notificationEventRepository}
}

func (n *NotificationEventAdapter) Save(event model.NotificationEvent) error {
	entity := mapper.ToNotificationEventEntity(&event)
	return n.notificationEventRepository.Save(entity)
}

func (n *NotificationEventAdapter) MarkAsProcessed(eventID string) error {
	notificationEvent, err := n.notificationEventRepository.GetByID(eventID)
	if err != nil {
		return err
	}
	notificationEvent.Processed = true
	return n.notificationEventRepository.Save(notificationEvent)
}

var _ spi.NotificationEventPersistence = (*NotificationEventAdapter)(nil)
