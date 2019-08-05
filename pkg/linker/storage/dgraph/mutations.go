package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo"
	dgraph_api "github.com/dgraph-io/dgo/protos/api"
	"github.com/pkg/errors"
)

func (c *Client) storeObject(ctx context.Context, txn *dgo.Txn, object interface{}) (string, error) {
	raw, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		return "", errors.Wrap(err, "encoding object")
	}

	fmt.Println("mutating:\n", string(raw))
	assigned, err := txn.Mutate(ctx, &dgraph_api.Mutation{SetJson: raw})
	if err != nil {
		return "", errors.Wrap(err, "running mutation")
	}

	uid, ok := assigned.Uids["blank-0"]
	if !ok {
		return "", errors.New("no resulting uid found")
	}

	return uid, nil
}
