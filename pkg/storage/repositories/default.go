package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/registries"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// Default repository for persisting api objects
// This is used for api resources that don't define custom repositories
type Default struct {
	Registry *storage.Registry
}

// NewDefault repository using the default registry
func NewDefault() storage.Repository {
	return &Default{
		Registry: registries.Default,
	}
}

// Add an object to the repository
func (r *Default) Add(ctx context.Context, resource storage.Resource, obj interface{}) (meta.ObjectMetaAccessor, error) {

	// resource, err := r.Registry.Get(obj.Kind())
	// if err != nil {
	// 	return nil, err
	// }

	if err := resource.BeforeCreation(ctx, obj); err != nil {
		return nil, err
	}

	// TODO: validation

	converted, err := resource.Converter.ToStorage(obj)
	if err != nil {
		return nil, err
	}

	log.From(ctx).Info("writing object", zap.String("converted", fmt.Sprintf("%#v", converted)))
	stored, err := resource.Provider.Write(converted)
	if err != nil {
		return nil, err
	}

	if stored == nil {
		return nil, errors.New("storing returned nil object")
	}

	new, ok := obj.(meta.ObjectMetaAccessor)
	if !ok {
		return nil, errors.New("invalid stored object")
	}

	return new, nil
}

// Get an object from the repository
func (r *Default) Get(ctx context.Context, resource storage.Resource, kind meta.GroupVersionKind, id string) (meta.ObjectMetaAccessor, error) {

	// resource, err := r.Registry.Get(kind)
	// if err != nil {
	// 	return nil, err
	// }

	log.From(ctx).Info("reading object", zap.Stringer("kind", kind), zap.String("id", id))
	stored, err := resource.Provider.Read(id)
	if err != nil {
		return nil, err
	}

	// TODO: updating?

	converted, err := resource.Converter.FromStorage(stored)
	if err != nil {
		return nil, err
	}

	obj, ok := converted.(meta.ObjectMetaAccessor)
	if !ok {
		return nil, fmt.Errorf("invalid stored object: %#v", converted)
	}

	return obj, nil
}
