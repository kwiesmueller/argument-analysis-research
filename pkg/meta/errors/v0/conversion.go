package errors

import (
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	apiError "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors"
	"github.com/pkg/errors"
)

// Conversion is an error indicating the requested object could not be converted
type Conversion struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *ConversionData  `json:"data"`
}

// Kind is an access helper to implement the KindAccessor interface
func (c *Conversion) Kind() meta.GroupVersionKind {
	return ConversionKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *Conversion) GetObjectMeta() meta.Object {
	return c.Metadata
}

// ConversionKind defines the object GroupVersionKind
var ConversionKind = apiError.Conversion.WithVersion(APIVersion)

// ConversionData describes the lookup parameters and detailed error message
type ConversionData struct {
	Reason  string                `json:"reason"`
	From    meta.GroupVersionKind `json:"from"`
	To      meta.GroupVersionKind `json:"to"`
	Message string                `json:"message"`
	Error   error                 `json:"error,omitempty"`
	Details string                `json:"details"`
}

// NewConversion initializes an empty object with the correct metadata
func NewConversion(reason string, from, to meta.GroupVersionKind, details string) *Conversion {
	err := errors.New("conversion failed")
	return &Conversion{
		Metadata: meta.NewObjectMeta(ConversionKind),
		Data: &ConversionData{
			Reason:  reason,
			From:    from,
			To:      to,
			Message: err.Error(),
			Error:   err,
			Details: details,
		},
	}
}

// Error implements the error interface
func (c *Conversion) Error() string {
	if c.Data == nil {
		return fmt.Sprintf("%s: undefined", c.Kind())
	}
	return fmt.Sprintf("%s: %s: (%s > %s): %s: %s", c.Kind(), c.Data.Error, c.Data.From, c.Data.To, c.Data.Reason, c.Data.Details)
}
