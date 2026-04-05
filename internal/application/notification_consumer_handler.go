package application

import (
	"encoding/json"

	"github.com/lian-corp/notification-microservice/internal/domain/api"
	"github.com/lian-corp/notification-microservice/internal/domain/model"
)

type NotificationConsumerHandler struct {
	notificationService api.NotificationService
}

func NewNotificationConsumerHandler(notificationService api.NotificationService) *NotificationConsumerHandler {
	return &NotificationConsumerHandler{
		notificationService: notificationService,
	}
}

func (h *NotificationConsumerHandler) HandleMessage(message []byte) error {
	var event model.NotificationEvent
	if err := json.Unmarshal([]byte(message), &event); err != nil {
		return err
	}
	return h.notificationService.ProcessEvent(event)
}
