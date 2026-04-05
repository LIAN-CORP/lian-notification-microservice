package model

import "time"

type DebtSnapshot struct {
	DebtID string
	ClientID string
	TotalAmount float64
	RemainingAmount float64
	UpdatedAt time.Time
}