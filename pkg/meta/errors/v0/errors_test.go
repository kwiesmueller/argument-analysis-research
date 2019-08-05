package errors_test

import (
	"testing"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors"
	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
)

func TestNotFoundKind(t *testing.T) {
	err := errorsv0.NewNotFound("test lookup")
	if !errors.IsKind(err, errorsv0.NotFoundKind) {
		t.Fatal("error kind does not match expectation")
	}

	testKind := meta.GroupVersionKind{
		Group:   meta.Group("tests"),
		Version: meta.Version("v0"),
		Kind:    meta.Kind("testKind"),
	}

	if errors.IsKind(err, testKind) {
		t.Fatal("error kind does not match expectation")
	}
}
