package meta

import "time"

// ObjectMetaAccessor for accessing object metadata
type ObjectMetaAccessor interface {
	GetObjectMeta() Object
}

// Object lets you work with object metadata
type Object interface {
	GetID() string
	SetID(string)
	GetGroupVersionKind() GroupVersionKind
	SetGroupVersionKind(GroupVersionKind)
	GetContext() *Context
	SetContext(*Context)
	GetCreated() time.Time
	SetCreated(time.Time)
	GetLabels() map[string]string
	SetLabels(map[string]string)
}

// Kind provides access to an objects GroupVersionKind
// type Kind interface {
// 	Kind() GroupVersionKind
// }
