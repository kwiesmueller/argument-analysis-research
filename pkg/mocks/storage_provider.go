package mocks

import (
	"sync"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// StorageProvider mocks the pkg/storage/Provider interface
type StorageProvider struct {
	sync.RWMutex
	Store map[string]interface{}
}

var _ storage.Provider = &StorageProvider{}

// NewStorageProvider for testing
func NewStorageProvider() *StorageProvider {
	return &StorageProvider{
		Store: make(map[string]interface{}),
	}
}

func (p *StorageProvider) Write(obj interface{}) (interface{}, error) {
	accessor, ok := obj.(storage.IDAccessor)
	if !ok {
		return nil, errors.New("object does not implement the storage.IDAccessor interface")
	}

	p.Lock()
	defer p.Unlock()
	p.Store[accessor.GetID()] = obj

	return obj, nil
}

func (p *StorageProvider) Read(id string) (interface{}, error) {
	p.RLock()
	defer p.RUnlock()

	obj, ok := p.Store[id]
	if !ok {
		return nil, errors.NewNotFound(id)
	}

	return obj, nil
}
