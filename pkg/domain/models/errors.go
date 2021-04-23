package models

import "errors"

var (
	// ErrMissingRequiredField error raised when missing a required field
	ErrMissingRequiredField = errors.New("missing required field")
	// ErrInvalidAmount error raised when the app receive a invalid amount
	ErrInvalidAmount = errors.New("invalid amount received")
)
