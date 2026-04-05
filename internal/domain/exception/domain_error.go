package exception

type DomainError struct {
	Message string
}

func (e DomainError) Error() string {
	return e.Message
}

func NewDomainError(message string) error {
	return DomainError{Message: message}
}