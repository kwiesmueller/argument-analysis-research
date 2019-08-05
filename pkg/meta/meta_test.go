package meta

import (
	"encoding/json"
	"testing"

	"k8s.io/utils/diff"
)

func TestMetaJSONEncoding(t *testing.T) {
	meta := &ObjectMeta{
		ID: "test",
		Kind: GroupVersionKind{
			Version: "vt1",
			Group:   "testing/meta",
			Kind:    "encoding",
		},
	}

	b, err := json.Marshal(meta)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	expected := []byte(`{"apiVersion":"vt1","kind":"testing/meta.encoding","id":"test","created":"0001-01-01T00:00:00Z"}`)

	if string(b) != string(expected) {
		t.Errorf("encoding meta to json failed,\nexpected:\n%s\ngot:\n%s\n", expected, b)
	}
}

func TestMetaJSONDecoding(t *testing.T) {
	data := []byte(`{"apiVersion":"vt1","kind":"testing/meta.encoding","id":"test"}`)

	expect := &ObjectMeta{
		ID: "test",
		Kind: GroupVersionKind{
			Version: "vt1",
			Group:   "testing/meta",
			Kind:    "encoding",
		},
	}

	meta := &ObjectMeta{}

	err := json.Unmarshal(data, &meta)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	d := diff.ObjectReflectDiff(expect, meta)
	if d != "<no diffs>" {
		t.Errorf("decoding meta from json failed, diff: %v", d)
	}
}
