package linking

import (
	"reflect"

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
	path := SegmentConversion.ToStorage()

	segment, typeOK := obj.(*linking.Segment)
	if !typeOK {
		return path.Fail("invalid type", "")
	}

	if _, err := path.ValidateAPIObject(segment); err != nil {
		return nil, err
	}

	if segment.Data == nil {
		return path.Fail("invalid object data", "")
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
	path := SegmentConversion.FromStorage()

	segment, typeOK := obj.(*Segment)
	if !typeOK {
		return path.Fail("invalid type", reflect.TypeOf(obj).String())
	}

	if segment == nil {
		return path.Fail("invalid object", "")
	}

	if _, err := path.CheckStorageObjectKind(segment); err != nil {
		return nil, err
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
