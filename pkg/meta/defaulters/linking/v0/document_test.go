package defaulters

import (
	"testing"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
)

func TestHandleRequestLinkerID(t *testing.T) {
	var tests = []struct {
		name   string
		doc    *linking.Document
		id     string
		expect bool
	}{
		{
			name:   "nil-''",
			doc:    &linking.Document{Data: &linking.DocumentData{}},
			id:     "",
			expect: true,
		},
		{
			name:   "{}-''",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{}}},
			id:     "",
			expect: true,
		},
		{
			name:   "{''}-''",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{""}}},
			id:     "",
			expect: true,
		},
		{
			name:   "{1}-''",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{"1"}}},
			id:     "",
			expect: false,
		},
		{
			name:   "nil-1",
			doc:    &linking.Document{Data: &linking.DocumentData{}},
			id:     "1",
			expect: true,
		},
		{
			name:   "{}-1",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{}}},
			id:     "1",
			expect: true,
		},
		{
			name:   "{''}-1",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{""}}},
			id:     "1",
			expect: false,
		},
		{
			name:   "{1}-1",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{"1"}}},
			id:     "1",
			expect: true,
		},
		{
			name:   "{1}-2",
			doc:    &linking.Document{Data: &linking.DocumentData{Linkers: []string{"1"}}},
			id:     "2",
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ok := handleRequestLinkerID(test.doc, test.id); ok != test.expect {
				t.Fatalf("unexpected result wanted:%v\ngot:%#v", test.expect, ok)
			}
		})
	}
}
