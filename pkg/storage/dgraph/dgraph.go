package dgraph

// type Storage interface {
// 	SetMetadata(ctx context.Context, meta *Metadata)
// 	Metadata(ctx context.Context) (*Metadata, error)

// 	InsertDocument(ctx context.Context, doc *document.Document)
// 	InsertSegment(ctx context.Context, doc *Segment)
// 	InsertLink(ctx context.Context, doc *Edge)

// 	Documents(ctx context.Context) (map[string]*document.Document, error)
// 	Segments(ctx context.Context) (map[string]*Segment, error)
// 	Links(ctx context.Context) (map[string]*Edge, error)
// }

// type StorageManager interface {
// 	New(ctx context.Context, id string) (Storage, error)
// 	List(ctx context.Context) (map[string]Storage, error)
// }

import (
	"context"

	"github.com/dgraph-io/dgo"
)

// Client provides a storage implementation based on https://dgraph.io to provide an actionable data layer for the discovered results
type Client struct {
	*dgo.Dgraph
}

// ObjectUID queries for an object with the given id and returns its UID
func (c *Client) ObjectUID(ctx context.Context, domain string) (string, error) {
	txn := c.Dgraph.NewReadOnlyTxn()
	defer txn.Commit(ctx)

	return c.queryUID(ctx, txn, domain)
}
