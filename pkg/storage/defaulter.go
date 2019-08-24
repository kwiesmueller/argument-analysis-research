package storage

import "context"

// Defaulter defines the interface for defaulting objects before handling them inside the system
type Defaulter interface {
	BeforeCreation(ctx context.Context, obj interface{}) error
	BeforeUpdate(ctx context.Context, obj interface{}) error
}
