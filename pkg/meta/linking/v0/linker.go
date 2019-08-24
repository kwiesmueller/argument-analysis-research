package linking

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// Linker for separating different datasets and interacting with them
type Linker struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *LinkerData      `json:"data"`
}

// Kind is an access helper to implement the KindAccessor interface
func (c *Linker) Kind() meta.GroupVersionKind {
	return LinkerKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *Linker) GetObjectMeta() meta.Object {
	return c.Metadata
}

// LinkerKind defines the object GroupVersionKind
var LinkerKind = Linking.WithKind("linker")

// LinkerData describes a linker instance
type LinkerData struct {
	UID         string  `json:"uid"`
	Description string  `json:"description,omitempty"`
	Rater       string  `json:"rater"`
	Threshold   float32 `json:"threshold"`
}

// NewLinker initializes an empty object with the correct metadata
func NewLinker(data *LinkerData) *Linker {
	if data == nil {
		data = &LinkerData{}
	}
	return &Linker{
		Metadata: meta.NewObjectMeta(LinkerKind),
		Data:     data,
	}
}
