package dgraph

import (
	"context"
	"encoding/json"

	errorsv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/errors/v0"

	"github.com/dgraph-io/dgo"
	"github.com/pkg/errors"
)

func (c *Client) queryUID(ctx context.Context, txn *dgo.Txn, query string) (string, error) {
	resp, err := txn.Query(ctx, query)
	if err != nil {
		return "", errors.Wrap(err, "querying for uid")
	}

	var result uidResult
	if err := json.Unmarshal(resp.GetJson(), &result); err != nil {
		return "", errors.Wrap(err, "decoding uid response")
	}

	if len(result.All) < 1 {
		return "", errorsv0.NewNotFound(query)
	}

	return result.All[0].UID, nil
}

func (c *Client) queryObject(ctx context.Context, txn *dgo.Txn, query string, obj interface{}) error {
	resp, err := txn.Query(ctx, query)
	if err != nil {
		return errors.Wrap(err, "querying for object")
	}

	if err := json.Unmarshal(resp.GetJson(), &obj); err != nil {
		return errors.Wrap(err, "decoding object response")
	}

	return nil
}
