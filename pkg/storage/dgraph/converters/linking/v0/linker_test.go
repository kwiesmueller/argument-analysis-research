package linking_test

import (
	"encoding/json"
	"testing"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	storage "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"

	fuzz "github.com/google/gofuzz"
)

func TestLinkerRoundTrip(t *testing.T) {

	linkerData := &linking.LinkerData{}
	linker := linking.NewLinker(linkerData)

	fuzzer := fuzz.New().NilChance(0)
	fuzzer.Fuzz(linkerData)
	fuzzer.Fuzz(linker)

	// TODO handle additional metadata
	linker.Metadata.Kind = linking.LinkerKind
	linker.Metadata.Labels = nil
	linker.Metadata.Context = nil

	linkerConverter := &storage.LinkerConverter{
		MetadataConverter: &storage.MetadataConverter{},
	}
	storageLinker, err := linkerConverter.ToStorage(linker)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := storageLinker.(*storage.Linker); !ok {
		t.Fatalf("invalid type from conversion: %#v", storageLinker)
	}

	roundTrippedLinker, err := linkerConverter.FromStorage(storageLinker)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := roundTrippedLinker.(*linking.Linker); !ok {
		t.Fatalf("invalid type from conversion: %#v", roundTrippedLinker)
	}

	rawLinker, err := json.Marshal(linker)
	if err != nil {
		t.Fatalf("json conversion: %v", err)
	}
	rawRTLinker, err := json.Marshal(roundTrippedLinker)
	if err != nil {
		t.Fatalf("json conversion: %v", err)
	}

	if string(rawLinker) != string(rawRTLinker) {
		t.Fatalf("roundtrip object mismatch:\nori:\n%s\nnew:\n%s", rawLinker, rawRTLinker)
	}
}

func TestNilToStorageWontPanic(t *testing.T) {

	linkerData := &linking.LinkerData{}
	linker := linking.NewLinker(linkerData)

	fuzzer := fuzz.New().NilChance(0.8)
	fuzzer.Fuzz(linkerData)
	fuzzer.Fuzz(linker)

	linkerConverter := &storage.LinkerConverter{
		MetadataConverter: &storage.MetadataConverter{},
	}
	storageLinker, err := linkerConverter.ToStorage(linker)
	if err == nil {
		t.Fatal(err)
	}

	if _, ok := storageLinker.(*storage.Linker); ok {
		t.Fatalf("unexpected valid type from conversion: %#v", storageLinker)
	}
}
func TestNilFromStorageWontPanic(t *testing.T) {

	linkerData := &linking.LinkerData{}
	linker := &storage.Linker{LinkerData: linkerData}

	fuzzer := fuzz.New().NilChance(0.8)
	fuzzer.Fuzz(linkerData)
	fuzzer.Fuzz(linker)

	linkerConverter := &storage.LinkerConverter{
		MetadataConverter: &storage.MetadataConverter{},
	}

	apiLinker, err := linkerConverter.FromStorage(linker)
	if err == nil {
		t.Fatal(err)
	}

	if _, ok := apiLinker.(*linking.Linker); ok {
		t.Fatalf("unexpected valid type from conversion: %#v", apiLinker)
	}
}
