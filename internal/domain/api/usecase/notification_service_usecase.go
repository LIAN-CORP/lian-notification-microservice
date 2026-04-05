package usecase

import (
	"github.com/lian-corp/notification-microservice/internal/domain/api"
	"github.com/lian-corp/notification-microservice/internal/domain/model"
)

type NotificationServiceUseCase struct {
	notificationEventService api.NotificationEventService
}

func NewNotificationServiceUseCase(notificationEventService api.NotificationEventService) *NotificationServiceUseCase {
	return &NotificationServiceUseCase{
		notificationEventService: notificationEventService,
	}
}

func (s *NotificationServiceUseCase) ProcessEvent(event model.NotificationEvent) error {
	return nil
}

var _ api.NotificationService = (*NotificationServiceUseCase)(nil)