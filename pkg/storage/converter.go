package storage

import (
	"fmt"
	"sync"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
)

// Converter defines the required interface for an object converter
type Converter interface {
	ToStorage(interface{}) (interface{}, error)
	FromStorage(interface{}) (interface{}, error)
}

// Conversion allows to declare conversion paths and provides helpers around error handling
type Conversion struct {
	// API is the api level GroupVersionKind of the resource
	API meta.GroupVersionKind `json:"api"`
	// Storage is the storage level GroupVersionKind of the resource
	Storage meta.GroupVersionKind `json:"storage"`

	onceTo, onceFrom       sync.Once
	toStorage, fromStorage *ConversionPath
}

// ToStorage returns a respectively directed ConversionPath
func (c *Conversion) ToStorage() *ConversionPath {
	c.onceTo.Do(func() { c.toStorage = &ConversionPath{Conversion: c, Direction: DirectionToStorage} })
	return c.toStorage
}

// FromStorage returns a respectively directed ConversionPath
func (c *Conversion) FromStorage() *ConversionPath {
	c.onceFrom.Do(func() { c.fromStorage = &ConversionPath{Conversion: c, Direction: DirectionFromStorage} })
	return c.fromStorage
}

// CheckObjectMeta ensures the passed in object is not nil and contains metadata
func (c *Conversion) CheckObjectMeta(obj meta.ObjectMetaAccessor) bool {
	if obj == nil {
		return false
	}
	if obj.GetObjectMeta().IsNil() {
		return false
	}
	return true
}

type conversionDirection string

const (
	// DirectionToStorage indicates an API->Storage ConversionPath
	DirectionToStorage = conversionDirection("toStorage")
	// DirectionFromStorage indicates an Storage->API ConversionPath
	DirectionFromStorage = conversionDirection("fromStorage")
)

// ConversionPath defines a specific from->to relationship based on a Conversion
type ConversionPath struct {
	*Conversion
	Direction conversionDirection `json:"direction"`
}

// Path returns the conversion path
func (c *ConversionPath) Path() (meta.GroupVersionKind, meta.GroupVersionKind) {
	if c.Direction == DirectionToStorage {
		return c.API, c.Storage
	}
	return c.Storage, c.API
}

// From returns the conversion origin
func (c *ConversionPath) From() meta.GroupVersionKind {
	if c.Direction == DirectionToStorage {
		return c.API
	}
	return c.Storage
}

// To returns the conversion origin
func (c *ConversionPath) To() meta.GroupVersionKind {
	if c.Direction == DirectionFromStorage {
		return c.API
	}
	return c.Storage
}

// CheckAPIObjectKind verifies the passed in api object's kind matches the ConversionPath's API GroupVersionKind
func (c *ConversionPath) CheckAPIObjectKind(obj meta.ObjectMetaAccessor) (interface{}, error) {
	return c.checkKind(obj.GetObjectMeta().GetGroupVersionKind())
}

// CheckStorageObjectKind verifies the passed in storage object's kind matches the ConversionPath's API GroupVersionKind
// TODO: do we really want to store objects with their api gvk?
func (c *ConversionPath) CheckStorageObjectKind(obj meta.KindAccessor) (interface{}, error) {
	return c.checkKind(obj.Kind())
}

// checkKind to be the internal API GroupVersionKind
func (c *ConversionPath) checkKind(kind meta.GroupVersionKind) (interface{}, error) {
	if kind.Is(c.API) {
		return nil, nil
	}
	return c.Fail("invalid object kind", fmt.Sprintf("got: %s", kind))
}

// ValidateAPIObject combines the required verification calls for an api object before converting it to storage
func (c *ConversionPath) ValidateAPIObject(obj meta.ObjectMetaAccessor) (interface{}, error) {
	if !c.CheckObjectMeta(obj) {
		return c.Fail("invalid object", "")
	}

	return c.CheckAPIObjectKind(obj)
}

// Fail the conversion providing a full return value set for a Converter function
// The returned error is a meta conversion error
func (c *ConversionPath) Fail(reason, details string) (interface{}, error) {
	from, to := c.Path()
	return nil, errorsv0.NewConversion(reason, from, to, details)
}
