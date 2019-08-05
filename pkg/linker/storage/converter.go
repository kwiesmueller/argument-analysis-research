package storage

// Converter defines the required interface for an object converter
type Converter interface {
	ToStorage(interface{}) (interface{}, error)
	FromStorage(interface{}) (interface{}, error)
}
