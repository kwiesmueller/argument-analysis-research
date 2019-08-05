package linking

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// Link between at least two segments for indicating their relation
type Link struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *LinkData        `json:"data"`
}

// Kind is an access helper to implement the Kind interface
func (c *Link) Kind() meta.GroupVersionKind {
	return LinkKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *Link) GetObjectMeta() meta.Object {
	return c.Metadata
}

// LinkKind defines the object GroupVersionKind
var LinkKind = Linking.WithKind("link")

// LinkData describes the link and its weight
type LinkData struct {
	UID    string     `json:"uid"`
	Weight float32    `json:"weight"`
	Links  []*Segment `json:"links"`
}

// NewLink initializes an empty object with the correct metadata
func NewLink(data *LinkData) *Link {
	return &Link{
		Metadata: meta.NewObjectMeta(LinkKind),
		Data:     &LinkData{},
	}
}
