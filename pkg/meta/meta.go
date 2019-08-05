package meta

import (
	"encoding/json"
	"time"
)

// ObjectMeta is metadata that all internal objects must have
type ObjectMeta struct {
	ID      string            `json:"id,omitempty"`
	Kind    GroupVersionKind  `json:"-"`
	Context *Context          `json:"context,omitempty"`
	Created time.Time         `json:"created,omitempty"`
	Labels  map[string]string `json:"labels,omitempty"`
}

// NewObjectMeta with kind returns a initialized metadata object with the creation timestamp set
func NewObjectMeta(kind GroupVersionKind) *ObjectMeta {
	return &ObjectMeta{
		Kind:    kind,
		Created: time.Now().UTC(),
	}
}

func (m *ObjectMeta) GetObjectMeta() Object                   { return m }
func (m *ObjectMeta) GetID() string                           { return m.ID }
func (m *ObjectMeta) SetID(to string)                         { m.ID = to }
func (m *ObjectMeta) GetGroupVersionKind() GroupVersionKind   { return m.Kind }
func (m *ObjectMeta) SetGroupVersionKind(to GroupVersionKind) { m.Kind = to }
func (m *ObjectMeta) GetContext() *Context                    { return m.Context }
func (m *ObjectMeta) SetContext(to *Context)                  { m.Context = to }
func (m *ObjectMeta) GetLabels() map[string]string            { return m.Labels }
func (m *ObjectMeta) SetLabels(to map[string]string)          { m.Labels = to }

// GetCreated returns the creation timestamp of an object in UTC
func (m *ObjectMeta) GetCreated() time.Time { return m.Created.UTC() }

// SetCreated sets the creation timestamp of an object in UTC
func (m *ObjectMeta) SetCreated(to time.Time) { m.Created = to.UTC() }

// MarshalJSON implements custom marshaling to encode the objects GVK into two fields, apiVersion and a GroupKind
// It encodes meta timestamps as UTC
func (m *ObjectMeta) MarshalJSON() ([]byte, error) {
	type Alias ObjectMeta
	type meta struct {
		APIVersion string `json:"apiVersion"`
		GroupKind  string `json:"kind"`
		*Alias
	}

	obj := meta{
		APIVersion: m.Kind.Version.String(),
		GroupKind:  m.Kind.GroupKind().String(),
		Alias:      (*Alias)(m),
	}
	obj.Alias.Created = obj.Alias.Created.UTC()

	return json.Marshal(obj)
}

// UnmarshalJSON implements custom unmarshalling to decode the objects apiVersion and kind into a full GroupVersionKind
// It decodes meta timestamps as UTC
func (m *ObjectMeta) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	type Alias ObjectMeta
	type meta struct {
		APIVersion string `json:"apiVersion"`
		GroupKind  string `json:"kind"`
		*Alias
	}

	obj := &meta{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(b, obj); err != nil {
		return err
	}

	m.Kind = ParseGroupKind(obj.GroupKind).WithVersion(Version(obj.APIVersion))
	m.Created = m.Created.UTC()

	return nil
}
