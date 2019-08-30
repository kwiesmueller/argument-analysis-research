package defaulters

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

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
		return errors.New("invalid object")
	}

	linkerID := chi.URLParamFromCtx(ctx, "id")
	if !handleRequestLinkerID(document, linkerID) {
		return errors.New("request linkerID does not match object")
	}
	if len(linkerID) < 1 {
		return errors.New("missing linkerID")
	}

	if len(document.Data.Content) < 1 {
		return errors.New("missing content")
	}

	h := sha256.New()
	h.Write([]byte(document.Data.Content))
	document.Metadata.ID = hex.EncodeToString(h.Sum(nil))

	for i, segment := range document.Data.Segments {
		if segment.Metadata.IsNil() {
			segment = linking.NewSegment(segment.Data)
			h := sha256.New()
			h.Write([]byte(segment.Data.Content))
			segment.Metadata.ID = hex.EncodeToString(h.Sum(nil))
			document.Data.Segments[i] = segment
		}
	}

	return nil
}

// BeforeUpdate document won't be touched
// TODO: set an updated timestamp?
func (m *Document) BeforeUpdate(_ context.Context, _ interface{}) error {
	return nil
}

func handleRequestLinkerID(doc *linking.Document, id string) bool {
	if doc.Data.Linkers == nil || len(doc.Data.Linkers) < 1 {
		doc.Data.Linkers = []string{id}
		return true
	}

	for _, linker := range doc.Data.Linkers {
		if linker == id {
			return true
		}
	}

	if len(doc.Data.Linkers) > 0 {
		return false
	}

	return false
}
