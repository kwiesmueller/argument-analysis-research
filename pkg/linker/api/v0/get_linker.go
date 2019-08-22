package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"

	"github.com/pkg/errors"
)

// GetLinker provides a handler for retrieving linker instances by their ID
func GetLinker(ctx context.Context, repo *storage.Repository) api.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		id := chi.URLParam(r, "id")
		if len(id) < 1 {
			http.Error(w, "missing id", http.StatusBadRequest)
			return errors.New("missing id")
		}

		old, err := repo.Get(ctx, linkingv0.LinkerKind, id)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(old); err != nil {
			return err
		}

		return nil
	}
}
