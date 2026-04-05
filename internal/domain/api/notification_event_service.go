package api

import "github.com/lian-corp/notification-microservice/internal/domain/model"

type NotificationEventService interface {
	CreateEvent(event model.NotificationEvent)
}