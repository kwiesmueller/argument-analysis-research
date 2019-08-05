package meta

// Context of an object to store various information passed on to other services
type Context struct {
	Trace []byte `json:"trace,omitempty"`
}
