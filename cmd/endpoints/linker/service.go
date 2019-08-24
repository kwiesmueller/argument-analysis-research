package main

import (
	"fmt"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/api"
	linker_api "github.com/canonical-debate-lab/argument-analysis-research/pkg/linker/api/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/service"

	linkingv0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/meta/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/mocks"
	linkingStoragev0 "github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/dgraph/converters/linking/v0"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/registries"
	"github.com/canonical-debate-lab/argument-analysis-research/pkg/storage/repositories"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
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

	// --- Mock Provider --- //

	provider := mocks.NewStorageProvider()
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
	provider.Store[linker.Metadata.ID] = linker

	repo := repositories.NewDefault()
	registries.PrepareDefault(provider, repo)

	// ---

	srv.Router.Post("/linker", api.NewHandler(ctx, linker_api.NewLinker(ctx, registries.Default)))
	srv.Router.Get("/linker/{id}", api.NewHandler(ctx, linker_api.GetLinker(ctx, registries.Default)))

	srv.Router.Post("/document", api.NewHandler(ctx, linker_api.NewDocument(ctx, registries.Default)))
	srv.Router.Post("/linker/{id}/document", api.NewHandler(ctx, linker_api.NewDocument(ctx, registries.Default)))

	if err := srv.Start(ctx); err != nil {
		log.From(ctx).Fatal("running server", zap.Error(err))
	}

	log.From(ctx).Info("finished")
}
