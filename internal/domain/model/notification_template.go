package model

import "github.com/lian-corp/notification-microservice/internal/domain/constants"

type NotificationTemplate struct {
	ID string
	EventType constants.EventType
	Template string
}