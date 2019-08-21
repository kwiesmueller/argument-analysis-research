package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"

	"github.com/pkg/errors"
)

// NewLinkerHandler for creating linker instances
func NewLinkerHandler(ctx context.Context) api.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) api.Response {
		var (
			req  *linkingv0.Linker
			resp = &Response{Error: &api.Error{}}
		)

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Fail(errors.Wrap(err, "decoding request"))
			return resp
		}

		return resp
	}
}
