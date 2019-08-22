package linking

import (
	"fmt"
	"strings"
	"time"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/meta"
	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"
)

// Metadata defines the fields shared between all linking storage objects
type Metadata struct {
	APIVersion string    `json:"apiVersion"`
	GroupKind  string    `json:"kind"`
	ID         string    `json:"id"`
	Created    time.Time `json:"created,omitempty"`
	Labels     []string  `json:"label,omitempty"`
}

// Kind is an access helper to implement the meta.KindAccessor interface
func (m Metadata) Kind() meta.GroupVersionKind {
	return meta.ParseGroupKind(m.GroupKind).WithVersion(meta.Version(m.APIVersion))
}

// GetID is an access helper to implement the storage.IDAccessort interface
func (m Metadata) GetID() string {
	return m.ID
}

var (
	// ObjectMetaKind for the storage subgroup
	ObjectMetaKind = meta.GroupVersionKind{
		Group:   meta.Group("storage.api"),
		Version: Linking.Version,
		Kind:    meta.ObjectMetaKind.Kind,
	}
	// MetadataConversion between the API and Storage kind
	MetadataConversion = storage.Conversion{
		API:     meta.ObjectMetaKind,
		Storage: ObjectMetaKind,
	}
)

// MetadataConverter for converting metadata to/from its dgraph storage representations
type MetadataConverter struct{}

// ToStorage converts a passed in linking/v0.metadata to it's storage format
func (c *MetadataConverter) ToStorage(obj *meta.ObjectMeta) (Metadata, error) {
	if obj == nil {
		return Metadata{}, errorsv0.NewConversion("invalid object",
			meta.ObjectMetaKind, ObjectMetaKind, "",
		)
	}

	var labels []string
	for key, value := range obj.Labels {
		labels = append(labels, fmt.Sprintf("%s:%s", key, value))
	}

	return Metadata{
		APIVersion: obj.Kind.Version.String(),
		GroupKind:  obj.Kind.GroupKind().String(),
		ID:         obj.ID,
		Created:    obj.Created,
		Labels:     labels,
	}, nil
}

// FromStorage converts a passed in metadata from storage format to a linking/v0.metadata
func (c *MetadataConverter) FromStorage(obj Metadata) (*meta.ObjectMeta, error) {
	kind := obj.Kind()
	out := meta.NewObjectMeta(kind)
	out.ID = obj.ID
	out.Created = obj.Created

	for _, label := range obj.Labels {
		parts := strings.Split(label, ":")
		if len(parts) < 2 {
			return nil, errorsv0.NewConversion("invalid label",
				meta.ObjectMetaKind, ObjectMetaKind, fmt.Sprintf("got: %s", label),
			)
		}

		if out.Labels == nil {
			out.Labels = make(map[string]string)
		}

		out.Labels[parts[0]] = strings.Join(parts[1:], ":")
	}

	return out, nil
}
