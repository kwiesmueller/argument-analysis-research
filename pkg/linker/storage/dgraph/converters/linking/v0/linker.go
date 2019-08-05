package linking

import (
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/storage"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
)

// Linker as the storage representation of linking/v0.linker
type Linker struct {
	BaseData
	*linking.LinkerData
}

// LinkerConverter for converting linking/v0.linker objects to/from their dgraph storage representations
type LinkerConverter struct{}

var _ storage.Converter = &LinkerConverter{}

// ToStorage converts a passed in linking/v0.linker to it's storage format
func (c *LinkerConverter) ToStorage(obj interface{}) (interface{}, error) {
	linker, typeOK := obj.(*linking.Linker)
	if !typeOK {
		return nil, errorsv0.NewConversion("invalid type",
			linking.LinkerKind.String(), "storage/linking/v0/linker", "",
		)
	}

	if linker == nil {
		return nil, errorsv0.NewConversion("invalid object",
			linking.LinkerKind.String(), "storage/linking/v0/linker", "",
		)
	}

	if linker.Kind() != linking.LinkerKind {
		return nil, errorsv0.NewConversion("invalid object kind",
			linking.LinkerKind.String(), "storage/linking/v0/linker",
			fmt.Sprintf("got: %s", linker.Kind()),
		)
	}

	return &Linker{
		BaseData: BaseData{
			APIVersion: linking.LinkerKind.Version.String(),
			Kind:       linking.LinkerKind.GroupKind().String(),
			ID:         linker.Metadata.ID,
			Created:    linker.Metadata.Created,
		},
		LinkerData: linker.Data,
	}, nil
}

// FromStorage converts a passed in linker from storage format to a linking/v0.linker
func (c *LinkerConverter) FromStorage(obj interface{}) (interface{}, error) {
	linker, typeOK := obj.(*Linker)
	if !typeOK {
		return nil, errorsv0.NewConversion("invalid type",
			"storage/linking/v0/linker", linking.LinkerKind.String(), "",
		)
	}

	if linker == nil {
		return nil, errorsv0.NewConversion("invalid object",
			"storage/linking/v0/linker", linking.LinkerKind.String(), "",
		)
	}

	kind := meta.ParseGroupKind(linker.Kind).WithVersion(meta.Version(linker.APIVersion))

	if kind != linking.LinkerKind {
		return nil, errorsv0.NewConversion("invalid object kind",
			"storage/linking/v0/linker", linking.LinkerKind.String(),
			fmt.Sprintf("got: %s", kind),
		)
	}

	out := linking.NewLinker(linker.LinkerData)
	out.Metadata.ID = linker.ID
	out.Metadata.Created = linker.Created

	return out, nil
}
