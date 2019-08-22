package storage

// IDAccessor provides a common interface to access storage object IDs
type IDAccessor interface {
	GetID() string
}
