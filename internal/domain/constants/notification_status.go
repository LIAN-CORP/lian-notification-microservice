package constants

type NotificationStatus string

const (
	StatusPending NotificationStatus = "PENDING"
	StatusSent NotificationStatus = "SENT"
	StatusFailed NotificationStatus = "FAILED"
)