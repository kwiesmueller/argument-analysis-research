package linking

import (
	"fmt"

	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// Segment as the storage representation of linking/v0.segment
type Segment struct {
	Metadata
	UID     string `json:"uid"`
	Content string `json:"content"`
}

var (
	// SegmentKind for the storage subgroup
	SegmentKind = Linking.WithKind(linking.SegmentKind.Kind)
	// SegmentConversion between the API and Storage kind
	SegmentConversion = storage.Conversion{
		API:     linking.SegmentKind,
		Storage: SegmentKind,
	}
)

// SegmentConverter for converting linking/v0.segment objects to/from their dgraph storage representations
type SegmentConverter struct {
	MetadataConverter *MetadataConverter
}

var _ storage.Converter = &SegmentConverter{}

// ToStorage converts a passed in linking/v0.segment to it's storage format
func (c *SegmentConverter) ToStorage(obj interface{}) (interface{}, error) {
	segment, typeOK := obj.(*linking.Segment)
	if !typeOK {
		return nil, errorsv0.NewConversion("invalid type",
			linking.SegmentKind, SegmentKind, "",
		)
	}

	if segment == nil || segment.Metadata == nil {
		return nil, errorsv0.NewConversion("invalid object",
			linking.SegmentKind, SegmentKind, "",
		)
	}

	if !segment.Kind().Is(linking.SegmentKind) {
		return nil, errorsv0.NewConversion("invalid object kind",
			linking.SegmentKind, SegmentKind,
			fmt.Sprintf("got: %s", segment.Kind()),
		)
	}

	if segment.Data == nil {
		return nil, errorsv0.NewConversion("invalid object data",
			linking.SegmentKind, SegmentKind,
			fmt.Sprintf("got: %s", segment.Data),
		)
	}

	metadata, err := c.MetadataConverter.ToStorage(segment.Metadata)
	if err != nil {
		return nil, err
	}

	return &Segment{
		Metadata: metadata,
		UID:      segment.Data.UID,
		Content:  segment.Data.Content,
	}, nil
}

// FromStorage converts a passed in segment from storage format to a linking/v0.segment
func (c *SegmentConverter) FromStorage(obj interface{}) (interface{}, error) {
	segment, typeOK := obj.(*Segment)
	if !typeOK {
		return nil, errorsv0.NewConversion("invalid type",
			SegmentKind, linking.SegmentKind, "",
		)
	}

	if segment == nil {
		return nil, errorsv0.NewConversion("invalid object",
			SegmentKind, linking.SegmentKind, "",
		)
	}

	kind := segment.Kind()

	if kind != linking.SegmentKind {
		return nil, errorsv0.NewConversion("invalid object kind",
			SegmentKind, linking.SegmentKind,
			fmt.Sprintf("got: %s", kind),
		)
	}

	metadata, err := c.MetadataConverter.FromStorage(segment.Metadata)
	if err != nil {
		return nil, err
	}

	out := &linking.Segment{
		Metadata: metadata,
		Data: &linking.SegmentData{
			UID:     segment.UID,
			Content: segment.Content,
		},
	}

	return out, nil
}
