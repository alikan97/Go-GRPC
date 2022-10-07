package model

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization      Type = "AUTHORIZATION"
	BadRequest         Type = "BAD_REQUEST"
	Conflict           Type = "CONFLICT"
	Internal           Type = "INTERNAL"
	NotFound           Type = "NOTFOUND"
	PayloadTooLarge    Type = "PAYLOAD_TOO_LARGE"
	ServiceUnavailable Type = "SERVICE_UNAVAILABLE"
)

type Error struct {
	Type    Type   `json:"type"` // json tags used to marshal/serialize json
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// Map errors to http response type
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason %v", reason),
	}
}
func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}
func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error...",
	}
}
func NewPayloadTooLarge(maxBodysize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodysize, contentLength),
	}
}
func NewServiceUnavailable() *Error {
	return &Error{
		Type:    ServiceUnavailable,
		Message: fmt.Sprint("Service unavailable"),
	}
}
