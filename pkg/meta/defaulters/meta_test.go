package defaulters_test

import (
	"context"
	"testing"

	meta_defaulters "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/defaulters"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
)

func TestMetadataDefaulting(t *testing.T) {
	ctx := context.Background()
	defaulter := &meta_defaulters.Metadata{}

	linker := linking.NewLinker(&linking.LinkerData{})

	if err := defaulter.BeforeCreation(ctx, linker); err != nil {
		t.Fatal(err)
	}

	id := linker.Metadata.ID
	if len(id) < 1 {
		t.Fatal("defaulting failed: id not set")
	}

	if err := defaulter.BeforeCreation(ctx, linker); err != nil {
		t.Fatal(err)
	}

	if id == linker.Metadata.ID {
		t.Fatalf("id did not change: %s -> %s", id, linker.Metadata.ID)
	}
}
