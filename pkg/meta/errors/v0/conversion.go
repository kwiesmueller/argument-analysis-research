package errors

import (
	"errors"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	apiError "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors"
)

// Conversion is an error indicating the requested object could not be converted
type Conversion struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *ConversionData  `json:"data"`
}

// Kind is an access helper to implement the Kind interface
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
	Reason  string `json:"reason"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message error  `json:"error"`
	Details string `json:"details"`
}

// NewConversion initializes an empty object with the correct metadata
func NewConversion(reason, from, to, details string) *Conversion {
	return &Conversion{
		Metadata: meta.NewObjectMeta(ConversionKind),
		Data: &ConversionData{
			Reason:  reason,
			From:    from,
			To:      to,
			Message: errors.New("conversion failed"),
			Details: details,
		},
	}
}

// Error implements the error interface
func (c *Conversion) Error() string {
	if c.Data == nil {
		return fmt.Sprintf("%s: undefined", c.Kind())
	}
	return fmt.Sprintf("%s: %s: (%s > %s): %s", c.Kind(), c.Data.Message, c.Data.From, c.Data.To, c.Data.Details)
}
