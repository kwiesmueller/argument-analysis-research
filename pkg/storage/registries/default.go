package registries

import (
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"

	meta_defaulters "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/defaulters"
	linking_defaulters "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/defaulters/linking/v0"
	linking_v0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	linking_storage_v0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"
	linking_preparators_v0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/preparators/linking/v0"
)

// Default registry mainly intended for testing
// May be removed or moved into a testing package in the near future
var Default = storage.NewRegistry()

// PrepareDefault adds all default resources to the Default registry providing them with a shared storage provider
func PrepareDefault(sharedProvider storage.Provider, sharedRepository storage.Repository) {
	metadataConverter := &linking_storage_v0.MetadataConverter{}

	linkerConverter := &linking_storage_v0.LinkerConverter{
		MetadataConverter: metadataConverter,
	}

	segmentConverter := &linking_storage_v0.SegmentConverter{
		MetadataConverter: metadataConverter,
	}

	Default.Add(linking_v0.LinkerKind, storage.Resource{
		Defaulters: []storage.Defaulter{
			&meta_defaulters.Metadata{},
			&linking_defaulters.Linker{},
		},
		Converter:  linkerConverter,
		Repository: sharedRepository,
		Provider:   sharedProvider,
	})

	Default.Add(linking_v0.SegmentKind, storage.Resource{
		Defaulters: []storage.Defaulter{
			&meta_defaulters.Metadata{},
		},
		Converter:  segmentConverter,
		Repository: sharedRepository,
		Provider:   sharedProvider,
	})

	Default.Add(linking_v0.DocumentKind, storage.Resource{
		Defaulters: []storage.Defaulter{
			&meta_defaulters.Metadata{},
			&linking_defaulters.Document{},
		},
		Converter: &linking_storage_v0.DocumentConverter{
			MetadataConverter: metadataConverter,
			LinkerConverter:   linkerConverter,
			SegmentConverter:  segmentConverter,
		},
		Preparator: &linking_preparators_v0.Document{
			Registry: Default,
		},
		Repository: sharedRepository,
		Provider:   sharedProvider,
	})
}
