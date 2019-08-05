package meta

import (
	"encoding/json"
	"time"
)

// Incomplete allows partially parsing unstructured data (like json) to access its metadata, before parsing the payload
type Incomplete struct {
	Metadata *ObjectMeta     `json:"metadata"`
	Data     json.RawMessage `json:"data,omitempty" datastore:"data,flatten"`
}

var _ ObjectMetaAccessor = &Incomplete{}
var _ Object = &Incomplete{}

// GetObjectMeta to implement the accessor interface
func (u *Incomplete) GetObjectMeta() Object {
	if u == nil || u.Metadata == nil {
		return &ObjectMeta{}
	}
	return u.Metadata
}

func (u *Incomplete) GetID() string   { return u.GetObjectMeta().GetID() }
func (u *Incomplete) SetID(to string) { u.GetObjectMeta().SetID(to) }
func (u *Incomplete) GetGroupVersionKind() GroupVersionKind {
	return u.GetObjectMeta().GetGroupVersionKind()
}
func (u *Incomplete) SetGroupVersionKind(to GroupVersionKind) {
	u.GetObjectMeta().SetGroupVersionKind(to)
}
func (u *Incomplete) GetContext() *Context   { return u.GetObjectMeta().GetContext() }
func (u *Incomplete) SetContext(to *Context) { u.GetObjectMeta().SetContext(to) }
func (u *Incomplete) GetLabels() map[string]string {
	return u.GetObjectMeta().GetLabels()
}
func (u *Incomplete) SetLabels(to map[string]string) { u.GetObjectMeta().SetLabels(to) }

// GetCreated returns the creation timestamp of an object in UTC
func (u *Incomplete) GetCreated() time.Time { return u.GetObjectMeta().GetCreated().UTC() }

// SetCreated sets the creation timestamp of an object in UTC
func (u *Incomplete) SetCreated(to time.Time) { u.GetObjectMeta().SetCreated(to.UTC()) }
