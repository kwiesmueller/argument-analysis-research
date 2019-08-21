package errors

import (
	"errors"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	apiError "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors"
)

// UnknownKind is an error indicating the requested object could not be converted
type UnknownKind struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *UnknownKindData `json:"data"`
}

// Kind is an access helper to implement the KindAccessor interface
func (c *UnknownKind) Kind() meta.GroupVersionKind {
	return UnknownKindKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *UnknownKind) GetObjectMeta() meta.Object {
	return c.Metadata
}

// UnknownKindKind defines the object GroupVersionKind
var UnknownKindKind = apiError.UnknownKind.WithVersion(APIVersion)

// UnknownKindData describes the lookup parameters and detailed error message
type UnknownKindData struct {
	Kind    meta.GroupVersionKind `json:"kind"`
	Message error                 `json:"error"`
	Details string                `json:"details"`
}

// NewUnknownKind initializes an empty object with the correct metadata
func NewUnknownKind(kind meta.GroupVersionKind, details string) *UnknownKind {
	return &UnknownKind{
		Metadata: meta.NewObjectMeta(UnknownKindKind),
		Data: &UnknownKindData{
			Kind:    kind,
			Message: errors.New("unknown kind"),
			Details: details,
		},
	}
}

// Error implements the error interface
func (c *UnknownKind) Error() string {
	if c.Data == nil {
		return fmt.Sprintf("%s: undefined", c.Kind())
	}
	return fmt.Sprintf("%s: %s: %s (%s)", c.Kind(), c.Data.Message, c.Data.Details, c.Data.Kind)
}
