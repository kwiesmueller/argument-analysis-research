package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage"

	"github.com/dgraph-io/dgo"
	dgraph_api "github.com/dgraph-io/dgo/protos/api"
	"github.com/pkg/errors"
)

// txnContext .
type txnContext struct{}

var txnKey = txnContext{}

// TxnFromContext retrieves a transaction from context or creates a new one
func TxnFromContext(ctx context.Context, db *dgo.Dgraph) *dgo.Txn {
	if txn, ok := ctx.Value(txnKey).(*dgo.Txn); ok {
		return txn
	}
	return db.NewTxn()
}

// TxnToContext stores a transaction in context
func TxnToContext(ctx context.Context, txn *dgo.Txn) context.Context {
	return context.WithValue(ctx, txnKey, txn)
}

// Provider for a dgraph connection
type Provider struct {
	*dgo.Dgraph
}

var _ storage.Provider = &Provider{}

// NewProvider based on the passed in dgraph client
func NewProvider(dgraph *dgo.Dgraph) *Provider {
	return &Provider{
		Dgraph: dgraph,
	}
}

func (p *Provider) Write(ctx context.Context, obj interface{}) (interface{}, error) {
	raw, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return nil, errors.Wrap(err, "encoding object")
	}

	txn := TxnFromContext(ctx, p.Dgraph)

	fmt.Println("mutating:\n", string(raw))
	assigned, err := txn.Mutate(ctx, &dgraph_api.Mutation{SetJson: raw})
	if err != nil {
		return nil, errors.Wrap(err, "running mutation")
	}

	_, ok := assigned.Uids["blank-0"]
	if !ok {
		return nil, errors.New("no resulting uid found")
	}

	if err := txn.Commit(ctx); err != nil {
		return nil, err
	}

	return obj, nil
}

func (p *Provider) Read(ctx context.Context, lookup string) (interface{}, error) {
	txn := TxnFromContext(ctx, p.Dgraph)

	resp, err := txn.Query(ctx, lookup)
	if err != nil {
		return nil, errors.Wrap(err, "querying for object")
	}

	js := resp.GetJson()

	fmt.Println(string(js))

	type objResult struct {
		Object []interface{} `json:"object"`
	}

	var obj objResult
	if err := json.Unmarshal(js, &obj); err != nil {
		return nil, errors.Wrap(err, "decoding object response")
	}

	if len(obj.Object) < 1 {
		return nil, errors.New("not found")
	}

	return obj.Object[0], nil
}
