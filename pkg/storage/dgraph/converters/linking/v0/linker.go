package linking

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// Linker as the storage representation of linking/v0.linker
type Linker struct {
	Metadata
	*linking.LinkerData
}

var (
	// LinkerKind for the storage subgroup
	LinkerKind = Linking.WithKind(linking.LinkerKind.Kind)
	// LinkerConversion between the API and Storage kind
	LinkerConversion = &storage.Conversion{
		API:     linking.LinkerKind,
		Storage: LinkerKind,
	}
)

// LinkerConverter for converting linking/v0.linker objects to/from their dgraph storage representations
type LinkerConverter struct {
	MetadataConverter *MetadataConverter
}

var _ storage.Converter = &LinkerConverter{}

// ToStorage converts a passed in linking/v0.linker to it's storage format
func (c *LinkerConverter) ToStorage(obj interface{}) (interface{}, error) {
	path := LinkerConversion.ToStorage()

	linker, typeOK := obj.(*linking.Linker)
	if !typeOK {
		return path.Fail("invalid type", "")
	}

	if _, err := path.ValidateAPIObject(linker); err != nil {
		return nil, err
	}

	metadata, err := c.MetadataConverter.ToStorage(linker.Metadata)
	if err != nil {
		return nil, err
	}

	return &Linker{
		Metadata:   metadata,
		LinkerData: linker.Data,
	}, nil
}

// FromStorage converts a passed in linker from storage format to a linking/v0.linker
func (c *LinkerConverter) FromStorage(obj interface{}) (interface{}, error) {
	path := LinkerConversion.FromStorage()

	linker, typeOK := obj.(*Linker)
	if !typeOK {
		return path.Fail("invalid type", "")
	}

	if linker == nil {
		return path.Fail("invalid object", "")
	}

	if _, err := path.CheckStorageObjectKind(linker); err != nil {
		return nil, err
	}

	metadata, err := c.MetadataConverter.FromStorage(linker.Metadata)
	if err != nil {
		return nil, err
	}

	out := &linking.Linker{
		Metadata: metadata,
		Data:     linker.LinkerData,
	}

	return out, nil
}
