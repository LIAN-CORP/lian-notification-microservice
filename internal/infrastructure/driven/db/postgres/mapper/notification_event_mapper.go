package mapper

import (
	"encoding/json"

	"github.com/lian-corp/notification-microservice/internal/domain/model"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/entity"
	"github.com/lian-corp/notification-microservice/internal/infrastructure/driven/db/postgres/entity/types"
)

func ToNotificationEventEntity(n *model.NotificationEvent) *entity.NotificationEventEntity {
	if n == nil {
		return nil
	}

	var payloadMap types.JSONB
	_ = json.Unmarshal(n.Payload, &payloadMap)

	return &entity.NotificationEventEntity{
		ID: n.ID,
		EventType: string(n.EventType),
		DebtID: n.DebtID,
		ClientID: n.ClientID,
		Payload: payloadMap,
		ReceivedAt: n.ReceivedAt,
		Processed: n.Processed,
	}
}