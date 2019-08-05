package linking

import "time"

// BaseData defines the fields shared between all linking storage objects
type BaseData struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	ID         string    `json:"id"`
	Created    time.Time `json:"created,omitempty"`
}
