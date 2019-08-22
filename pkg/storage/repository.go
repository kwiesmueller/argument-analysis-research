package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
)

// Provider defines the interface a persistence layer implementation has to expose for handling api objects
type Provider interface {
	Write(interface{}) (interface{}, error)
	Read(id string) (interface{}, error)
}

// Repository for persisting api objects
type Repository struct {
	Registry *Registry
	Provider Provider
}

// Add an object to the repository
func (r *Repository) Add(ctx context.Context, obj meta.KindAccessor) (meta.ObjectMetaAccessor, error) {

	converter, err := r.Registry.Get(obj.Kind())
	if err != nil {
		return nil, err
	}

	// TODO: defaulting

	converted, err := converter.ToStorage(obj)
	if err != nil {
		return nil, err
	}

	log.From(ctx).Info("writing object", zap.String("converted", fmt.Sprintf("%#v", converted)))
	stored, err := r.Provider.Write(converted)
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
func (r *Repository) Get(ctx context.Context, kind meta.GroupVersionKind, id string) (meta.ObjectMetaAccessor, error) {

	converter, err := r.Registry.Get(kind)
	if err != nil {
		return nil, err
	}

	log.From(ctx).Info("reading object", zap.Stringer("kind", kind), zap.String("id", id))
	stored, err := r.Provider.Read(id)
	if err != nil {
		return nil, err
	}

	// TODO: defaulting

	converted, err := converter.FromStorage(stored)
	if err != nil {
		return nil, err
	}

	obj, ok := converted.(meta.ObjectMetaAccessor)
	if !ok {
		return nil, fmt.Errorf("invalid stored object: %#v", converted)
	}

	return obj, nil
}
