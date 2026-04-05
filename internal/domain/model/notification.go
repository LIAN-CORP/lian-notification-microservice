package model

import (
	"time"

	"github.com/lian-corp/notification-microservice/internal/domain/constants"
)

type Notification struct {
	ID string
	EvenetID string
	ClientID string
	PhoneNumber string
	Message string
	Status constants.NotificationStatus
	Provider string
	ProviderResponse string
	CreatedAt time.Time
	SentAt time.Time
}

func (n *Notification) MarkAsSent(response string) {
	n.Status = constants.StatusSent
	n.ProviderResponse = response
	n.SentAt = time.Now()
}

func (n *Notification) MarkAsFailed(response string) {
	n.Status = constants.StatusFailed
	n.ProviderResponse = response
	n.SentAt = time.Now()
}