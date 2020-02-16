package storage

import "context"

// Provider defines the interface a persistence layer implementation has to expose for handling api objects
type Provider interface {
	Write(context.Context, interface{}) (interface{}, error)
	Read(context.Context, string) (interface{}, error)
}
