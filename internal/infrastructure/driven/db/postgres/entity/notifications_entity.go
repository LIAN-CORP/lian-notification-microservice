package entity

import (
	"time"

	"github.com/google/uuid"
)

type NotificationsEntity struct {
	ID 					uuid.UUID 				`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventID 			uuid.UUID 				`gorm:"type:uuid;index"`
	Event 				NotificationEventEntity `gorm:"foreignKey:EventID;references:ID"`
	ClientID 			uuid.UUID 				`gorm:"type:uuid;index"`
	PhoneNumber 		string 					`gorm:"type:varchar(20)"`
	Message 			string 					`gorm:"type:text"`
	Status 				string 					`gorm:"type:varchar(20);default:'PENDING';index"`
	Provider 			string 					`gorm:"type:varchar(50);default:'AWS_SNS'"`
	ProviderResponse 	string 					`gorm:"type:text"`
	CreatedAt 			time.Time 				`gorm:"type:timestamp;default:now()"`
	SentAt 				time.Time 				`gorm:"type:timestamp;default:now()"`
}