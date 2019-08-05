package errors

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// APIError defines the Group for errors
const APIError = meta.Group("errors")

// NotFound errors indicate the requested object or resource was not found
var NotFound = meta.GroupKind{
	Group: APIError,
	Kind:  "notFound",
}

// Conversion errors indicate a failed conversion between two types or version
var Conversion = meta.GroupKind{
	Group: APIError,
	Kind:  "conversion",
}

// UnknownKind errors indicate an objects kind is unknown to the api or a registry
var UnknownKind = meta.GroupKind{
	Group: APIError,
	Kind:  "unknownKind",
}

// IsKind checks if the passed in error is an api object and if it matches the provided GroupVersionKind
func IsKind(err error, kind meta.GroupVersionKind) bool {

	obj, err := meta.Accessor(err)
	if err != nil {
		return false
	}

	if obj.GetGroupVersionKind() != kind {
		return false
	}

	return true
}
