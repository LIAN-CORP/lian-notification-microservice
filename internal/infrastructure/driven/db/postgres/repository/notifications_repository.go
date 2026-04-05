package repository

import (
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/entity"
	"gorm.io/gorm"
)

type NotificationsRepository struct{
	db *gorm.DB
}

func NewNotificationsRepository(db *gorm.DB) *NotificationsRepository {
	return &NotificationsRepository{db: db}
}

func (r *NotificationsRepository) CreateNotification(notification *entity.NotificationsEntity) error {
	return r.db.Create(notification).Error
}

func (r *NotificationsRepository) GetNotifications() ([]entity.NotificationsEntity, error) {
	var notifications []entity.NotificationsEntity
	if err := r.db.Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}