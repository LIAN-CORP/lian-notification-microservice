package spi

type SMSPersistence interface {
	Send(phone string, message string) (string, error)
}