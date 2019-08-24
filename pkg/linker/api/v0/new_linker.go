package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"

	"github.com/pkg/errors"
)

// NewLinker provides a handler for creating linker instances
// POST /linker
func NewLinker(ctx context.Context, registry *storage.Registry) api.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		var obj *linkingv0.Linker

		err := json.NewDecoder(r.Body).Decode(&obj)
		if err != nil {
			return errors.Wrap(err, "decoding request")
		}

		resource, err := registry.Get(obj.Kind())
		if err != nil {
			return err
		}

		new, err := resource.Repository.Add(ctx, resource, obj)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(new); err != nil {
			return err
		}

		return nil
	}
}
