package main

import (
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	linker_api "github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/api/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/service"

	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	linkingStoragev0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/registries"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/repositories"

	"github.com/dgraph-io/dgo"
	dgraph_api "github.com/dgraph-io/dgo/protos/api"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Spec for the service
type Spec struct {
	service.BaseSpec
}

func main() {
	var svc Spec
	ctx := service.Init("linker-api", &svc)
	defer service.Defer(ctx)

	srv := api.New(fmt.Sprintf(":%d", svc.Port), svc.Debug)
	go srv.GracefulHandler(ctx)

	dial, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.From(ctx).Fatal("creating grpc client", zap.Error(err))
	}

	log.From(ctx).Debug("connecting dgraph")
	client := dgo.NewDgraphClient(dgraph_api.NewDgraphClient(dial))
	txn := client.NewReadOnlyTxn()
	if _, err := txn.Query(ctx, `schema {}`); err != nil {
		log.From(ctx).Fatal("validating dgraph connection", zap.Error(err))
	}

	provider := dgraph.NewProvider(client)

	// --- Mock Provider --- //

	linker := &linkingStoragev0.Linker{
		Metadata: linkingStoragev0.Metadata{
			APIVersion: linkingv0.LinkerKind.Version.String(),
			GroupKind:  linkingv0.LinkerKind.GroupKind().String(),
			ID:         "test",
		},
		LinkerData: &linkingv0.LinkerData{
			Description: "test",
			Rater:       "https://test.democracy.ovh/rater",
			Threshold:   0.25,
		},
	}
	if _, err := provider.Write(ctx, linker); err != nil {
		log.From(ctx).Fatal("inserting test linker", zap.Error(err))
	}

	repo := repositories.NewDefault()
	registries.PrepareDefault(provider, repo)

	// ---

	srv.Router.Post("/linker", api.NewHandler(ctx, linker_api.NewLinker(ctx, registries.Default)))
	srv.Router.Get("/linker/{id}", api.NewHandler(ctx, linker_api.GetLinker(ctx, registries.Default)))

	srv.Router.Post("/document", api.NewHandler(ctx, linker_api.NewDocument(ctx, registries.Default)))
	srv.Router.Post("/linker/{id}/document", api.NewHandler(ctx, linker_api.NewDocument(ctx, registries.Default)))
	srv.Router.Get("/document/{id}", api.NewHandler(ctx, linker_api.GetDocument(ctx, registries.Default)))

	if err := srv.Start(ctx); err != nil {
		log.From(ctx).Fatal("running server", zap.Error(err))
	}

	log.From(ctx).Info("finished")
}
