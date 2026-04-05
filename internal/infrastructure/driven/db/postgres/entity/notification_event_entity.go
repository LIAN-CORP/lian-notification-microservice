package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/entity/types"
)

type NotificationEventEntity struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventType 	string 		`gorm:"type:varchar(255)"`
	DebtID 		uuid.UUID 	`gorm:"type:uuid;index"`
	ClientID 	uuid.UUID 	`gorm:"type:uuid;index"`
	Payload 	types.JSONB	`gorm:"type:jsonb"`
	ReceivedAt 	time.Time 	`gorm:"type:timestamp;default:now()"`
	Processed 	bool 		`gorm:"type:boolean;default:false;index"`
}