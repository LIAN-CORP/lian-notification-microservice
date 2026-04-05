package usecase

import (
	"github.com/lian-corp/notification-microservice/internal/domain/api"
	"github.com/lian-corp/notification-microservice/internal/domain/model"
	"github.com/lian-corp/notification-microservice/internal/domain/spi"
)

type NotificationEventUseCase struct {
	notificationEventPersistence spi.NotificationEventPersistence
}

func NewNotificationEventUseCase(eventPersistence spi.NotificationEventPersistence) *NotificationEventUseCase {
	return &NotificationEventUseCase{notificationEventPersistence: eventPersistence}
}

func (e *NotificationEventUseCase) CreateEvent(event model.NotificationEvent) {
	e.notificationEventPersistence.Save(event)
}

var _ api.NotificationEventService = (*NotificationEventUseCase)(nil)