package dgraph

import (
	"context"
	"fmt"

	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"

	"github.com/dgraph-io/dgo"
	"github.com/pkg/errors"
)

type uidResult struct {
	All []struct {
		UID string `json:"uid"`
	} `json:"all"`
}

// TODO: use QueryWithVars
const objectUIDForIDQuery = `
	{
		all(func: eq(id, "%v")) {
			uid
		}
	} 
`

func (c *Client) queryUIDForID(ctx context.Context, txn *dgo.Txn, id string) (string, error) {
	if len(id) < 1 {
		return "", errors.New("can not query object without id")
	}

	query := fmt.Sprintf(objectUIDForIDQuery, id)

	return c.queryUID(ctx, txn, query)
}

type linkerResponse struct {
	Linkers []*storedLinker `json:"getLinker"`
}

type storedLinker struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	linkingv0.LinkerData
}

const linkerObjectQuery = `
	{
		getLinker(func: eq(id, "%v")) {
			uid
			id
			description
			rater
			threshold
		}
	}
`
