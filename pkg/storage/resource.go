package storage

import (
	"context"

	"github.com/pkg/errors"
)

// Resource provides access to the required components to handle an objects storage operations
type Resource struct {
	Defaulters []Defaulter
	Preparator Preparator
	Converter  Converter
	Repository Repository
	Provider   Provider
}

// Validate the resource contains all required fields
func (r Resource) Validate() error {
	if r.Converter == nil {
		return errors.New("resource missing converter")
	}
	if r.Provider == nil {
		return errors.New("resource missing provider")
	}
	if r.Repository == nil {
		return errors.New("resource missing repository")
	}

	return nil
}

// BeforeCreation runs the defaulter chain on the passed in object, returning the first error
func (r Resource) BeforeCreation(ctx context.Context, obj interface{}) error {
	for _, defaulter := range r.Defaulters {
		if err := defaulter.BeforeCreation(ctx, obj); err != nil {
			return err
		}
	}

	if r.Preparator != nil {
		return r.Preparator.BeforeCreation(ctx, r, obj)
	}

	return nil
}

// BeforeUpdate runs the defaulter chain on the passed in object, returning the first error
func (r Resource) BeforeUpdate(ctx context.Context, obj interface{}) error {
	for _, defaulter := range r.Defaulters {
		if err := defaulter.BeforeUpdate(ctx, obj); err != nil {
			return err
		}
	}

	if r.Preparator != nil {
		return r.Preparator.BeforeUpdate(ctx, r, obj)
	}

	return nil
}
