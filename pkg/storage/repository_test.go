package storage_test

import (
	"context"
	"fmt"
	"testing"

	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/mocks"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
	linkingStoragev0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/registries"
)

func TestRepositoryAdd(t *testing.T) {
	provider := mocks.NewStorageProvider()

	repository := &storage.Repository{
		Registry: registries.Default,
		Provider: provider,
	}

	linker := linkingv0.NewLinker(&linkingv0.LinkerData{
		Description: "test",
		Rater:       "https://test.democracy.ovh/rater",
		Threshold:   0.25,
	})

	ctx := context.Background()
	new, err := repository.Add(ctx, linker)
	if err != nil {
		fmt.Println(new)
		t.Fatal(err)
	}
}

func TestRepositoryGet(t *testing.T) {
	provider := mocks.NewStorageProvider()

	repository := &storage.Repository{
		Registry: registries.Default,
		Provider: provider,
	}

	linker := &linkingStoragev0.Linker{
		Metadata: linkingStoragev0.Metadata{
			APIVersion: linkingv0.LinkerKind.Version.String(),
			GroupKind:  linkingv0.LinkerKind.GroupKind().String(),
			ID:         "test",
		},
		LinkerData: &linkingv0.LinkerData{
			Description: "test",
			Rater:       "https://test.democracy.ovh/rater",
			Threshold:   0.25,
		},
	}

	provider.Store[linker.Metadata.ID] = linker

	ctx := context.Background()
	old, err := repository.Get(ctx, linkingv0.LinkerKind, "test")
	if err != nil {
		fmt.Println(old)
		t.Fatal(err)
	}
}
