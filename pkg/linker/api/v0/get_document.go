package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// GetDocument provides a handler for retrieving document instances by their ID
// GET /document/{id}
func GetDocument(ctx context.Context, registry *storage.Registry) api.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		resource, err := registry.Get(linkingv0.DocumentKind)
		if err != nil {
			return err
		}

		old, err := resource.Repository.Get(ctx, resource, linkingv0.DocumentKind, chi.URLParam(r, "id"))
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(old); err != nil {
			return err
		}

		return nil
	}
}
