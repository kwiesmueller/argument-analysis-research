package linking_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	storage "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"

	fuzz "github.com/google/gofuzz"
	"k8s.io/utils/diff"
)

func TestDocumentRoundTrip(t *testing.T) {

	documentData := &linking.DocumentData{}
	document := linking.NewDocument(documentData)

	fuzzer := fuzz.New().NilChance(0).Funcs(
		func(m *meta.ObjectMeta, c fuzz.Continue) {
			c.Fuzz(&m)

			if m.Labels != nil {
				m.Labels = make(map[string]string)
				for i := c.RandUint64(); i < 0; i-- {
					key := strings.ReplaceAll(c.RandString(), ":", ".")
					value := c.RandString()
					m.Labels[key] = value
				}
			}
		},
		func(m *linking.Linker, c fuzz.Continue) {
			c.Fuzz(&m.Metadata)
			c.Fuzz(&m.Data)
			if m.Metadata == nil {
				m.Metadata = &meta.ObjectMeta{}
			}
			m.Metadata.Kind = linking.LinkerKind
			m.Metadata.Context = nil
		},
		func(m *linking.Document, c fuzz.Continue) {
			c.Fuzz(&m.Metadata)
			c.Fuzz(&m.Data)
			if m.Metadata == nil {
				m.Metadata = &meta.ObjectMeta{}
			}
			m.Metadata.Kind = linking.DocumentKind
			m.Metadata.Context = nil
		},
		func(m *linking.Segment, c fuzz.Continue) {
			c.Fuzz(&m.Metadata)
			c.Fuzz(&m.Data)
			if m.Metadata == nil {
				m.Metadata = &meta.ObjectMeta{}
			}
			m.Metadata.Kind = linking.SegmentKind
			m.Metadata.Context = nil
		},
	)
	fuzzer.Fuzz(documentData)
	fuzzer.Fuzz(document)

	// TODO handle additional metadata
	document.Metadata.Context = nil

	metadataConverter := &storage.MetadataConverter{}
	documentConverter := &storage.DocumentConverter{
		MetadataConverter: metadataConverter,
		LinkerConverter: &storage.LinkerConverter{
			MetadataConverter: metadataConverter,
		},
		SegmentConverter: &storage.SegmentConverter{
			MetadataConverter: metadataConverter,
		},
	}
	storageDocument, err := documentConverter.ToStorage(document)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(storageDocument)

	if _, ok := storageDocument.(*storage.Document); !ok {
		t.Fatalf("invalid type from conversion: %#v", storageDocument)
	}

	roundTrippedDocument, err := documentConverter.FromStorage(storageDocument)
	if err != nil {
		t.Fatal(err)
	}

	rtDocument, ok := roundTrippedDocument.(*linking.Document)
	if !ok {
		t.Fatalf("invalid type from conversion: %#v", roundTrippedDocument)
	}

	d := diff.ObjectReflectDiff(document, rtDocument)
	if d != "<no diffs>" {
		t.Fatalf("decoding object to incomplete failed, diff: %v", d)
	}
	rawDocument, err := json.Marshal(document)
	if err != nil {
		t.Fatalf("json conversion: %v", err)
	}
	rawRTDocument, err := json.Marshal(roundTrippedDocument)
	if err != nil {
		t.Fatalf("json conversion: %v", err)
	}

	if string(rawDocument) != string(rawRTDocument) {
		t.Fatalf("roundtrip object mismatch:\nori:\n%s\nnew:\n%s", rawDocument, rawRTDocument)
	}
}
