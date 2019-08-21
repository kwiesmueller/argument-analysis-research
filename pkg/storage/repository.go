package storage

// Repository for persisting api objects
type Repository struct {
	Registry *Registry
}

var DefaultRepository = &Repository{
	Registry: nil,
}
