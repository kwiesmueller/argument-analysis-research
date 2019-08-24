package errors

import (
	"errors"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	apiError "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors"
)

// InvalidResource is an error indicating the requested object could not be converted
type InvalidResource struct {
	Metadata *meta.ObjectMeta     `json:"metadata"`
	Data     *InvalidResourceData `json:"data"`
}

// Kind is an access helper to implement the KindAccessor interface
func (c *InvalidResource) Kind() meta.GroupVersionKind {
	return InvalidResourceKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *InvalidResource) GetObjectMeta() meta.Object {
	return c.Metadata
}

// InvalidResourceKind defines the object GroupVersionKind
var InvalidResourceKind = apiError.InvalidResource.WithVersion(APIVersion)

// InvalidResourceData describes the lookup parameters and detailed error message
type InvalidResourceData struct {
	Kind    meta.GroupVersionKind `json:"kind"`
	Message string                `json:"message"`
	Error   error                 `json:"error,omitempty"`
	Details string                `json:"details"`
}

// NewInvalidResource initializes an empty object with the correct metadata
func NewInvalidResource(kind meta.GroupVersionKind, details string) *InvalidResource {
	err := errors.New("invalid resource")
	return &InvalidResource{
		Metadata: meta.NewObjectMeta(InvalidResourceKind),
		Data: &InvalidResourceData{
			Kind:    kind,
			Error:   err,
			Message: err.Error(),
			Details: details,
		},
	}
}

// Error implements the error interface
func (c *InvalidResource) Error() string {
	if c.Data == nil {
		return fmt.Sprintf("%s: undefined", c.Kind())
	}
	return fmt.Sprintf("%s: %s: %s (%s)", c.Kind(), c.Data.Message, c.Data.Details, c.Data.Kind)
}
