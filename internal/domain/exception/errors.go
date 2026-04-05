package exception

var (
	ErrTemplateNotFound = NewDomainError("Template not found")
	ErrSMSSendFailed = NewDomainError("Failed to send SMS")
	ErrInvalidEvent = NewDomainError("Invalid event")
)