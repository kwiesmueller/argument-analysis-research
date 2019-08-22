package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// HandlerFunc represents a http.HandlerFunc returning an error
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

// NewHandler returns a standard http.HandlerFunc providing context for the HandlerFunc
func NewHandler(ctx context.Context, f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(r.Context(), w, r)
		if err != nil {
			log.From(ctx).Error("handling", zap.Error(err))

			if apiError, ok := err.(meta.ObjectMetaAccessor); ok {
				encErr := json.NewEncoder(w).Encode(apiError)
				if encErr == nil {
					return
				}
				log.From(ctx).Error("encoding api error", zap.Error(encErr))
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
