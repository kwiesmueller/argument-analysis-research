package linking

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
)

// Segment created from a documents content
type Segment struct {
	Metadata *meta.ObjectMeta `json:"metadata"`
	Data     *SegmentData     `json:"data"`
}

// Kind is an access helper to implement the KindAccessor interface
func (c *Segment) Kind() meta.GroupVersionKind {
	return SegmentKind
}

// GetObjectMeta is an access helper to implement the ObjectMetaAccessor interface
func (c *Segment) GetObjectMeta() meta.Object {
	return c.Metadata
}

// SegmentKind defines the object GroupVersionKind
var SegmentKind = Linking.WithKind("segment")

// SegmentData describing the segment in relation to its object
type SegmentData struct {
	UID     string `json:"uid"`
	Content string `json:"content"`
}

// NewSegment initializes an empty object with the correct metadata
func NewSegment(data *SegmentData) *Segment {
	if data == nil {
		data = &SegmentData{}
	}
	return &Segment{
		Metadata: meta.NewObjectMeta(SegmentKind),
		Data:     data,
	}
}
