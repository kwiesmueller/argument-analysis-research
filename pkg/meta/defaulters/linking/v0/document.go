package defaulters

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Document defaults an object's document
type Document struct{}

// BeforeCreation sets the linker id if provided on the request context
func (m *Document) BeforeCreation(ctx context.Context, obj interface{}) error {
	document, ok := obj.(*linking.Document)
	if !ok {
		return errors.New("invalid type")
	}

	if document.Data == nil {
		document.Data = &linking.DocumentData{}
	}

	linkerID := chi.URLParamFromCtx(ctx, "id")
	if len(linkerID) > 0 {
		if document.Data.Linker == nil {
			document.Data.Linker = linking.NewLinker(nil)
		}

		linkerMeta := document.Data.Linker.GetObjectMeta()
		if linkerMeta == nil {
			return errors.New("object contains invalid linker")
		}

		linkerMetaID := linkerMeta.GetID()

		if len(linkerMetaID) > 0 && linkerMetaID != linkerID {
			return errors.New("request linkerID does not match object")
		}

		linkerMeta.SetID(linkerID)
	}

	return nil
}

// BeforeUpdate document won't be touched
// TODO: set an updated timestamp?
func (m *Document) BeforeUpdate(_ context.Context, _ interface{}) error {
	return nil
}
