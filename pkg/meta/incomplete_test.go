package meta

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/gofuzz"
	"k8s.io/utils/diff"
)

type testObject struct {
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	Data     testData   `json:"data,omitempty"`
}

type testData struct {
	TestString string  `json:"testString,omitempty"`
	TestInt    int     `json:"testInt,omitempty"`
	TestBool   bool    `json:"testBool,omitempty"`
	TestObj    testObj `json:"testObj,omitempty"`
}

type testObj struct {
	TestString string `json:"testString,omitempty"`
	TestInt    int    `json:"testInt,omitempty"`
	TestBool   bool   `json:"testBool,omitempty"`
}

func TestIncompleteJSONRoundtrip(t *testing.T) {
	fuzzer := fuzz.New().Funcs(
		func(e *ObjectMeta, c fuzz.Continue) {
			c.Fuzz(&e)
			e.Kind.Group = strings.Replace(e.Kind.Group, ".", "", -1)
			e.Kind.Version = strings.Replace(e.Kind.Version, ".", "", -1)
			e.Kind.Kind = strings.Replace(e.Kind.Kind, ".", "", -1)
		},
	)

	N := 1000
	for i := 0; i < N; i++ {
		metadata := &ObjectMeta{}
		fuzzer.Fuzz(metadata)

		data := testData{}
		fuzzer.Fuzz(&data)

		object := testObject{
			Metadata: metadata,
			Data:     data,
		}

		b, err := json.Marshal(&object)
		if err != nil {
			t.Fatalf("encoding object to json failed: %v", err)
		}

		u := &Incomplete{}
		err = json.Unmarshal(b, &u)
		if err != nil {
			t.Fatalf("decoding object to incomplete failed: %v", err)
		}

		d := diff.ObjectReflectDiff(object.Metadata, u.GetObjectMeta())
		if d != "<no diffs>" {
			t.Fatalf("decoding object to incomplete failed, diff: %v", d)
		}

		b, err = json.Marshal(&u)
		if err != nil {
			t.Fatalf("encoding object to json failed: %v", err)
		}

		newObject := testObject{}
		err = json.Unmarshal(b, &newObject)
		if err != nil {
			t.Fatalf("decoding object to typed failed: %v", err)
		}

		d = diff.ObjectReflectDiff(object, newObject)
		if d != "<no diffs>" {
			t.Errorf("decoding object to typed failed, diff: %v", d)
		}
	}
}
