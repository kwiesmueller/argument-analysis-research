package repositories

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/registries"
)

// Default .
var Default = &storage.Repository{
	Registry: registries.Default,
}
