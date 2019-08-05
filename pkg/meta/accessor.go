package meta

import (
	"github.com/pkg/errors"
)

// ErrNotObject is returned when an object does not implement the Object interfaces
var ErrNotObject = errors.New("object does not implement the Object interfaces")

// Accessor takes an arbitrary object pointer and returns meta.Interface.
// obj must be a pointer to an API type. An error is returned if the minimum
// required fields are missing. Fields that are not required return the default
// value and are a no-op if set.
func Accessor(obj interface{}) (Object, error) {
	switch t := obj.(type) {
	case Object:
		return t, nil
	case ObjectMetaAccessor:
		if m := t.GetObjectMeta(); m != nil {
			return m, nil
		}
		return nil, ErrNotObject
	default:
		return nil, ErrNotObject
	}
}

// Matcher takes two arbitrary object pointers, gets their accessors and verifies they match in kind
func Matcher(a, b interface{}) (bool, error) {
	ameta, err := Accessor(a)
	if err != nil {
		return false, errors.Wrap(err, "accessing object a")
	}

	bmeta, err := Accessor(b)
	if err != nil {
		return false, errors.Wrap(err, "accessing object b")
	}

	if ameta.GetGroupVersionKind() != bmeta.GetGroupVersionKind() {
		return false, nil
	}

	return true, nil
}
