package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lian-corp/notification-microservice/internal/domain/constants"
)

type NotificationEvent struct {
	ID uuid.UUID
	EventType constants.EventType
	DebtID uuid.UUID
	ClientID uuid.UUID
	Payload json.RawMessage
	ReceivedAt time.Time
	Processed bool
}