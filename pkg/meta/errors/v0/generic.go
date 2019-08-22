package errors

import (
	"github.com/pkg/errors"
)

// New wraps the errors package
func New(msg string) error {
	return errors.New(msg)
}
