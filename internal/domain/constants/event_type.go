package constants

type EventType string

const (
	EventDebtCreated EventType = "DEBT_CREATED"
	EventDebtReminder EventType = "DEBT_REMINDER"
	EventDebtIncrease EventType = "DEBT_INCREASE"
	EventPaymentMade EventType = "PAYMENT_MADE"
	EventDebtFullPaid EventType = "DEBT_FULL_PAID"
)