package meta_test

import (
	"testing"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

func TestGroupEquality(t *testing.T) {
	gvkA := meta.GroupVersionKind{Group: "A", Version: "1", Kind: "a"}
	gvkB := meta.GroupVersionKind{Group: "B", Version: "1", Kind: "b"}
	gvkC := meta.GroupVersionKind{Group: "A", Version: "1", Kind: "a"}

	if gvkA.Is(gvkB) {
		t.Fatalf("gvks should not be equal: %s-%s", gvkA, gvkB)
	}

	if !gvkA.Is(gvkC) {
		t.Fatalf("gvks should be equal: %s-%s", gvkA, gvkC)
	}
}
