package common

import "errors"

type ServiceError struct {
	Code  int
	Error error
}

var (
	BadParameters   = NewServiceMessage(400, "bad parameters")
	Unauthorized    = NewServiceMessage(401, "access denied")
	PaymentRequired = NewServiceMessage(402, "payment required")
	Forbbiden       = NewServiceMessage(403, "Forbidden")
	NotFound        = NewServiceMessage(404, "not found")
	NotAcceptable   = NewServiceMessage(406, "not accpetable")
)

func NewServiceMessage(code int, message string) *ServiceError {
	return &ServiceError{Code: code, Error: errors.New(message)}
}
func NewServiceError(code int, err error) *ServiceError {
	return &ServiceError{Code: code, Error: err}
}
