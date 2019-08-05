package linking

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// Document is the source object for linking arguments, it contains the original content and metadata
type Document struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *DocumentData    `json:"data"`
}

// Kind is an access helper to implement the Kind interface
func (c *Document) Kind() meta.GroupVersionKind {
	return DocumentKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *Document) GetObjectMeta() meta.Object {
	return c.Metadata
}

// DocumentKind defines the object GroupVersionKind
var DocumentKind = Linking.WithKind("document")

// DocumentData describes the document, its content and links to a certain Linker
type DocumentData struct {
	UID      string     `json:"uid"`
	Linker   *Linker    `json:"linker,omitempty"`
	Content  string     `json:"content"`
	Segments []*Segment `json:"segments,omitempty"`
}

// NewDocument initializes an empty object with the correct metadata
func NewDocument(data *DocumentData) *Document {
	return &Document{
		Metadata: meta.NewObjectMeta(DocumentKind),
		Data:     &DocumentData{},
	}
}
