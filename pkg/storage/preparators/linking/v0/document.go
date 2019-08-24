package preparators

import (
	"context"

	linking_v0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
	"github.com/pkg/errors"
)

// Document preparator populating documents based on their references
type Document struct {
	Registry *storage.Registry
}

// BeforeCreation looks up the linker referenced in the document and populates it
func (r *Document) BeforeCreation(ctx context.Context, resource storage.Resource, obj interface{}) error {

	document, _ := obj.(*linking_v0.Document)
	// TODO: check ok?
	if document.Data != nil && document.Data.Linker != nil {
		linker := document.Data.Linker
		if !linker.GetObjectMeta().IsNil() {
			linkerResource, err := r.Registry.Get(linker.Kind())
			if err != nil {
				return errors.Wrap(err, "reading linker resource")
			}

			storedLinker, err := resource.Repository.Get(ctx, linkerResource, linker.Kind(), linker.GetObjectMeta().GetID())
			if err != nil {
				return errors.Wrap(err, "reading linker reference")
			}
			document.Data.Linker = storedLinker.(*linking_v0.Linker)
		}
	}

	return nil
}

// BeforeUpdate does nothing
func (r *Document) BeforeUpdate(ctx context.Context, resource storage.Resource, obj interface{}) error {
	return nil
}
