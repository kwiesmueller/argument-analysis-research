package errors

import (
	"errors"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	apiError "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors"
)

// NotFound is an error indicating the requested object was not found
type NotFound struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *NotFoundData    `json:"data"`
}

// Kind is an access helper to implement the KindAccessor interface
func (c *NotFound) Kind() meta.GroupVersionKind {
	return NotFoundKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *NotFound) GetObjectMeta() meta.Object {
	return c.Metadata
}

// NotFoundKind defines the object GroupVersionKind
var NotFoundKind = apiError.NotFound.WithVersion(APIVersion)

// NotFoundData describes the lookup parameters and detailed error message
type NotFoundData struct {
	LookUp  string `json:"lookup"`
	Message error  `json:"error"`
}

// NewNotFound initializes an empty object with the correct metadata
func NewNotFound(lookup string) *NotFound {
	return &NotFound{
		Metadata: meta.NewObjectMeta(NotFoundKind),
		Data: &NotFoundData{
			LookUp:  lookup,
			Message: errors.New("object not found"),
		},
	}
}

// Error implements the error interface
func (c *NotFound) Error() string {
	if c.Data == nil {
		return fmt.Sprintf("%s: undefined", c.Kind())
	}
	return fmt.Sprintf("%s: %s (%s)", c.Kind(), c.Data.Message, c.Data.LookUp)
}
