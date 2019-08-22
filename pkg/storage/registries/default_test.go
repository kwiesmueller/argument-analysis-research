package registries_test

import (
	"testing"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/registries"

	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
)

func TestDefaultRegistry(t *testing.T) {
	registry := registries.Default

	converter, err := registry.Get(linkingv0.LinkerKind)
	if err != nil {
		t.Fatal(err)
	}
	if converter == nil {
		t.Fatal("registry did not know kind")
	}
}
