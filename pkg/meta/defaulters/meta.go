package defaulters

import (
	"context"
	"time"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

// Metadata defaults an object's metadata
type Metadata struct{}

// BeforeCreation every object requires a unique id and a creation timestamp
func (m *Metadata) BeforeCreation(_ context.Context, obj interface{}) error {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return err
	}

	rand, err := uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "generating uuid")
	}

	id := rand.String()
	if len(id) < 1 {
		return errors.New("invalid uuid generated")
	}

	accessor.SetID(id)
	accessor.SetCreated(time.Now().UTC())

	return nil
}

// BeforeUpdate metadata won't be touched
// TODO: set an updated timestamp?
func (m *Metadata) BeforeUpdate(_ context.Context, _ interface{}) error {
	return nil
}
