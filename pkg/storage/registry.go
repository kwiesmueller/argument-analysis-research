package storage

import (
	"sync"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
)

// Registry for mapping api types to their respective converters
type Registry struct {
	m          sync.RWMutex
	converters map[meta.GroupVersionKind]Converter
}

// NewRegistry for mapping api types to their respective converters
func NewRegistry() *Registry {
	return &Registry{
		converters: make(map[meta.GroupVersionKind]Converter),
	}
}

// Add a converter for the given GroupVersionKind
func (r *Registry) Add(gvk meta.GroupVersionKind, converter Converter) {
	r.m.Lock()
	defer r.m.Unlock()
	r.converters[gvk] = converter
}

// Get a converter for the given GroupVersionKind
func (r *Registry) Get(gvk meta.GroupVersionKind) (Converter, error) {
	r.m.RLock()
	defer r.m.RUnlock()
	if converter, ok := r.converters[gvk]; ok {
		return converter, nil
	}
	return nil, errorsv0.NewUnknownKind(gvk, "no converter found")
}
