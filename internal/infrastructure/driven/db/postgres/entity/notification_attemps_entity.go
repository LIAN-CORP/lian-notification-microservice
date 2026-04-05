package entity

import (
	"time"

	"github.com/google/uuid"
)

type NotificationAttemptEntity struct {
	ID 				uuid.UUID 			`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NotificationID 	uuid.UUID 			`gorm:"type:uuid;index"`
	Notifications 	NotificationsEntity `gorm:"foreignKey:NotificationID;references:ID"`
	Status 			string 				`gorm:"type:varchar(20)"`
	ErrorMessage 	string 				`gorm:"type:text"`
	AttemptAt 		time.Time 			`gorm:"type:timestamp;default:now()"`
}