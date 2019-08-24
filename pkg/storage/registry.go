package storage

import (
	"sync"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
)

// Registry for mapping api types to their respective resources
type Registry struct {
	m         sync.RWMutex
	resources map[meta.GroupVersionKind]Resource
}

// NewRegistry for mapping api types to their respective resources
func NewRegistry() *Registry {
	return &Registry{
		resources: make(map[meta.GroupVersionKind]Resource),
	}
}

// Add a converter for the given GroupVersionKind
func (r *Registry) Add(gvk meta.GroupVersionKind, resource Resource) {
	r.m.Lock()
	defer r.m.Unlock()
	r.resources[gvk] = resource
}

// Get a resource for the given GroupVersionKind
func (r *Registry) Get(gvk meta.GroupVersionKind) (Resource, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	resource, ok := r.resources[gvk]
	if !ok {
		// TODO: pointer?
		return Resource{}, errorsv0.NewUnknownKind(gvk, "no resource found")
	}

	if err := resource.Validate(); err != nil {
		return Resource{}, errorsv0.NewInvalidResource(gvk, err.Error())
	}

	return resource, nil
}
