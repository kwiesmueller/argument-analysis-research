package linking

import (
	"fmt"

	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
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
	document, typeOK := obj.(*linking.Document)
	if !typeOK {
		return nil, errorsv0.NewConversion("invalid type",
			linking.DocumentKind, DocumentKind, "",
		)
	}

	if document == nil || document.Metadata == nil {
		return nil, errorsv0.NewConversion("invalid object",
			linking.DocumentKind, DocumentKind, "",
		)
	}

	if !document.Kind().Is(linking.DocumentKind) {
		return nil, errorsv0.NewConversion("invalid object kind",
			linking.DocumentKind, DocumentKind,
			fmt.Sprintf("got: %s", document.Kind()),
		)
	}

	if document.Data == nil {
		return nil, errorsv0.NewConversion("invalid object data",
			linking.DocumentKind, DocumentKind,
			fmt.Sprintf("got: %v", document.Data),
		)
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
	document, typeOK := obj.(*Document)
	if !typeOK {
		return nil, errorsv0.NewConversion("invalid type",
			DocumentKind, linking.DocumentKind, "",
		)
	}

	if document == nil {
		return nil, errorsv0.NewConversion("invalid object",
			DocumentKind, linking.DocumentKind, "",
		)
	}

	kind := document.Kind()

	if !kind.Is(linking.DocumentKind) {
		return nil, errorsv0.NewConversion("invalid object kind",
			DocumentKind, linking.DocumentKind,
			fmt.Sprintf("got: %s,", kind),
		)
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
