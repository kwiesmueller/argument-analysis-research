package storage

import "context"

// Preparator defines the interface to prepare objects for storage
// Preparation is different from defaulting as it allows for repository operations to fill up the object based on it's references
type Preparator interface {
	BeforeCreation(ctx context.Context, resource Resource, obj interface{}) error
	BeforeUpdate(ctx context.Context, resource Resource, obj interface{}) error
}
