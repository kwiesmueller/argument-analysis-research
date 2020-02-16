package storage

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// Repository defines the interface for persisting and retrieving api objects
type Repository interface {
	Add(context.Context, Resource, interface{}) (meta.ObjectMetaAccessor, error)
	Get(context.Context, Resource, meta.GroupVersionKind, string) (meta.ObjectMetaAccessor, error)
}
