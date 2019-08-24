package storage

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// Provider defines the interface a persistence layer implementation has to expose for handling api objects
type Provider interface {
	Write(interface{}) (interface{}, error)
	Read(string) (interface{}, error)
}

// Repository defines the interface for persisting api objects
type Repository interface {
	Add(context.Context, Resource, interface{}) (meta.ObjectMetaAccessor, error)
	Get(context.Context, Resource, meta.GroupVersionKind, string) (meta.ObjectMetaAccessor, error)
}
