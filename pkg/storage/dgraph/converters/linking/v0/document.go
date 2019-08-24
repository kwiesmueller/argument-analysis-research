package linking

import (
	"reflect"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// Document as the storage representation of linking/v0.document
type Document struct {
	Metadata
	UID      string     `json:"uid"`
	Linker   *Linker    `json:"linker,omitempty"`
	Content  string     `json:"content"`
	Segments []*Segment `json:"segment,omitempty"`
}

var (
	// DocumentKind for the storage subgroup
	DocumentKind = Linking.WithKind(linking.DocumentKind.Kind)
	// DocumentConversion between the API and Storage kind
	DocumentConversion = storage.Conversion{
		API:     linking.DocumentKind,
		Storage: DocumentKind,
	}
)

// DocumentConverter for converting linking/v0.document objects to/from their dgraph storage representations
type DocumentConverter struct {
	MetadataConverter *MetadataConverter
	LinkerConverter   *LinkerConverter
	SegmentConverter  *SegmentConverter
}

var _ storage.Converter = &DocumentConverter{}

// ToStorage converts a passed in linking/v0.document to it's storage format
func (c *DocumentConverter) ToStorage(obj interface{}) (interface{}, error) {
	path := DocumentConversion.ToStorage()

	document, typeOK := obj.(*linking.Document)
	if !typeOK {
		return path.Fail("invalid type", reflect.TypeOf(obj).String())
	}

	if _, err := path.ValidateAPIObject(document); err != nil {
		return nil, err
	}

	if document.Data == nil {
		return path.Fail("invalid object data", "")
	}

	linker, err := c.LinkerConverter.ToStorage(document.Data.Linker)
	if err != nil {
		return nil, err
	}

	var segments []*Segment
	for _, segment := range document.Data.Segments {
		convertedSegment, err := c.SegmentConverter.ToStorage(segment)
		if err != nil {
			return nil, err
		}

		segments = append(segments, convertedSegment.(*Segment))
	}

	metadata, err := c.MetadataConverter.ToStorage(document.Metadata)
	if err != nil {
		return nil, err
	}

	return &Document{
		Metadata: metadata,
		UID:      document.Data.UID,
		Linker:   linker.(*Linker),
		Content:  document.Data.Content,
		Segments: segments,
	}, nil
}

// FromStorage converts a passed in document from storage format to a linking/v0.document
func (c *DocumentConverter) FromStorage(obj interface{}) (interface{}, error) {
	path := DocumentConversion.FromStorage()

	document, typeOK := obj.(*Document)
	if !typeOK {
		return path.Fail("invalid type", reflect.TypeOf(obj).String())
	}

	if document == nil {
		return path.Fail("invalid object", "")
	}

	if _, err := path.CheckStorageObjectKind(document); err != nil {
		return nil, err
	}

	linker, err := c.LinkerConverter.FromStorage(document.Linker)
	if err != nil {
		return nil, err
	}

	var segments []*linking.Segment
	for _, segment := range document.Segments {
		convertedSegment, err := c.SegmentConverter.FromStorage(segment)
		if err != nil {
			return nil, err
		}

		segments = append(segments, convertedSegment.(*linking.Segment))
	}

	metadata, err := c.MetadataConverter.FromStorage(document.Metadata)
	if err != nil {
		return nil, err
	}

	out := &linking.Document{
		Metadata: metadata,
		Data: &linking.DocumentData{
			UID:      document.UID,
			Linker:   linker.(*linking.Linker),
			Content:  document.Content,
			Segments: segments,
		},
	}

	return out, nil
}
