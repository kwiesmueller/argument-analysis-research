package storage_test

import (
	"sync"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

type MockProvider struct {
	sync.RWMutex
	Store map[string]interface{}
}

var _ storage.Provider = &MockProvider{}

// NewMockProvider for testing
func NewMockProvider() *MockProvider {
	return &MockProvider{
		Store: make(map[string]interface{}),
	}
}

func (p *MockProvider) Write(obj interface{}) (interface{}, error) {
	accessor, ok := obj.(storage.IDAccessor)
	if !ok {
		return nil, errors.New("object does not implement the storage.IDAccessor interface")
	}

	p.Lock()
	defer p.Unlock()
	p.Store[accessor.GetID()] = obj

	return obj, nil
}

func (p *MockProvider) Read(id string) (interface{}, error) {
	p.RLock()
	defer p.RUnlock()

	obj, ok := p.Store[id]
	if !ok {
		return nil, errors.NewNotFound(id)
	}

	return obj, nil
}
