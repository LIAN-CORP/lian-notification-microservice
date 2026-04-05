package repository

import (
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/entity"
	"gorm.io/gorm"
)

type NotificationEventRepository struct{
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *NotificationEventRepository {
	return &NotificationEventRepository{db: db}
}

func (r *NotificationEventRepository) Save(event *entity.NotificationEventEntity) error {
	return r.db.Create(event).Error
}

func (r *NotificationEventRepository) GetByID(eventID string) (*entity.NotificationEventEntity, error) {
	var entity entity.NotificationEventEntity
	if err := r.db.First(&entity, "id = ?", eventID).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}