package defaulters

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"

	"github.com/pkg/errors"
)

// Linker defaults an object's linker
type Linker struct{}

// BeforeCreation set the default rater if non is provided
func (m *Linker) BeforeCreation(_ context.Context, obj interface{}) error {
	linker, ok := obj.(*linking.Linker)
	if !ok {
		return errors.New("invalid type")
	}

	if linker.Data == nil {
		linker.Data = &linking.LinkerData{}
	}

	if len(linker.Data.Rater) < 1 {
		linker.Data.Rater = "https://research.democracy.ovh/argument/adw"
	}

	return nil
}

// BeforeUpdate linker won't be touched
// TODO: set an updated timestamp?
func (m *Linker) BeforeUpdate(_ context.Context, _ interface{}) error {
	return nil
}
