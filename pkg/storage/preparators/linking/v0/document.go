package preparators

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// Document preparator populating documents based on their references
type Document struct {
	Registry *storage.Registry
}

// BeforeCreation looks up the linker referenced in the document and populates it
func (r *Document) BeforeCreation(ctx context.Context, resource storage.Resource, obj interface{}) error {
	return nil
}

// BeforeUpdate does nothing
func (r *Document) BeforeUpdate(ctx context.Context, resource storage.Resource, obj interface{}) error {
	return nil
}
