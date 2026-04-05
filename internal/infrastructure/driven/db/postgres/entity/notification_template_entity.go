package entity

import "github.com/google/uuid"

type NotificationTemplateEntity struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventType 	string 		`gorm:"type:varchar(255)"`
	Template 	string 		`gorm:"type:text"`
}