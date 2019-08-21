package registries

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"

	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	linkingStoragev0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"
)

// DefaultRegistry .
var DefaultRegistry = storage.NewRegistry()

// init .
func init() {
	DefaultRegistry.Add(linkingv0.LinkerKind, &linkingStoragev0.LinkerConverter{})
}
