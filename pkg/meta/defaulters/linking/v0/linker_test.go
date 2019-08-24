package defaulters_test

import (
	"context"
	"testing"

	linking_defaulters "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/defaulters/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
)

func TestLinkerDefaulting(t *testing.T) {
	ctx := context.Background()
	defaulter := &linking_defaulters.Linker{}

	linker := linking.NewLinker(&linking.LinkerData{})

	if err := defaulter.BeforeCreation(ctx, linker); err != nil {
		t.Fatal(err)
	}

	if linker.Data.Rater != "https://research.democracy.ovh/argument/adw" {
		t.Fatal("defaulting failed: rater got", linker.Data.Rater)
	}
}
