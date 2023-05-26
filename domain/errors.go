package domain

import "errors"

var (
	ErrEmailExists = errors.New("the email already exists")
)

type ResponseError struct {
	Message string `json:"message"`
}
