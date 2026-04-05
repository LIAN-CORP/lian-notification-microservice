package entity

import (
	"time"

	"github.com/google/uuid"
)

type DebtSnapshotEntity struct {
	DebtID 			uuid.UUID 	`gorm:"type:uuid;primaryKey;index"`
	ClientID 		uuid.UUID 	`gorm:"type:uuid;index"`
	TotalAmount 	float64 	`gorm:"type:numeric(15,2)"`
	RemainingAmount float64 	`gorm:"type:numeric(15,2)"`
	updatedAt 		time.Time 	`gorm:"type:timestamp;default:now()"`
}