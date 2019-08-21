package meta

import (
	"fmt"
	"strings"
)

// Group allows to summarize a collection of kinds with a certain version
type Group string

func (g Group) String() string {
	return string(g)
}

// Kind defines an object or resource kind
type Kind string

func (k Kind) String() string {
	return string(k)
}

// Version defines an object or resource version
type Version string

func (v Version) String() string {
	return string(v)
}

// GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion
// to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling
type GroupVersionKind struct {
	Group   Group   `json:"group,omitempty"`
	Version Version `json:"version,omitempty"`
	Kind    Kind    `json:"kind,omitempty"`
}

// Empty returns true if group, version, and kind are empty
func (gvk GroupVersionKind) Empty() bool {
	return len(gvk.Group) == 0 && len(gvk.Version) == 0 && len(gvk.Kind) == 0
}

// GroupKind returns a new GroupKind from the objects group and kind
func (gvk GroupVersionKind) GroupKind() GroupKind {
	return GroupKind{Group: gvk.Group, Kind: gvk.Kind}
}

// GroupVersion returns a GroupVersion from the objects group and version
func (gvk GroupVersionKind) GroupVersion() GroupVersion {
	return GroupVersion{Group: gvk.Group, Version: gvk.Version}
}

// Is returns true if the passed in gvk's values are equal to it's own
func (gvk GroupVersionKind) Is(to GroupVersionKind) bool {
	if gvk.Group != to.Group ||
		gvk.Version != to.Version ||
		gvk.Kind != to.Kind {
		return false
	}
	return true
}

func (gvk GroupVersionKind) String() string {
	return fmt.Sprintf("%s/%s.%s", gvk.Group, gvk.Version, gvk.Kind)
}

// GroupKind specifies a Group and a Kind, but does not force a version.  This is useful for identifying
// concepts during lookup stages without having partially valid types
type GroupKind struct {
	Group Group `json:"group,omitempty"`
	Kind  Kind  `json:"kind,omitempty"`
}

// Empty returns true if group and kind are empty
func (gk GroupKind) Empty() bool {
	return len(gk.Group) == 0 && len(gk.Kind) == 0
}

// WithVersion builds a GroupVersionKind based on the object and the passed in version
func (gk GroupKind) WithVersion(version Version) GroupVersionKind {
	return GroupVersionKind{Group: gk.Group, Version: version, Kind: gk.Kind}
}

func (gk GroupKind) String() string {
	if len(gk.Group) == 0 {
		return gk.Kind.String()
	}
	return fmt.Sprintf("%s.%s", gk.Group, gk.Kind)
}

// ParseGroupKind takes a string and decodes it into a GroupKind
func ParseGroupKind(from string) GroupKind {
	s := strings.Split(from, ".")
	if len(s) < 1 {
		return GroupKind{}
	}
	if len(s) < 2 {
		return GroupKind{Kind: Kind(s[0])}
	}
	return GroupKind{Group: Group(s[0]), Kind: Kind(s[1])}
}

// GroupVersion contains the "group" and the "version", which uniquely identifies the API.
type GroupVersion struct {
	Group   Group   `json:"group,omitempty"`
	Version Version `json:"version,omitempty"`
}

// Empty returns true if group and version are empty
func (gv GroupVersion) Empty() bool {
	return len(gv.Group) == 0 && len(gv.Version) == 0
}

// String puts "group" and "version" into a single "group/version" string.
func (gv GroupVersion) String() string {
	// special case the internal apiVersion for the legacy kube types
	if gv.Empty() {
		return ""
	}

	if len(gv.Group) > 0 {
		return fmt.Sprintf("%s/%s", gv.Group, gv.Version)
	}
	return gv.Version.String()
}

// WithKind creates a GroupVersionKind based on the method receiver's GroupVersion and the passed Kind.
func (gv GroupVersion) WithKind(kind Kind) GroupVersionKind {
	return GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
}
