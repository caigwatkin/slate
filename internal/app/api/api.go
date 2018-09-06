package api

import (
	"context"
	"net/http"
	"slate/internal/app/api/middleware"
	"slate/internal/app/api/route"
	"slate/internal/pkg/errors"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secret"

	"github.com/go-chi/chi"
)

type APIClient struct {
	config Config
	deps   Deps
	router *chi.Mux
}

type Config struct {
	Env          string
	GCPProjectID string
	Port         string
}

type Deps struct {
	SecretClient *secret.Client
}

func NewClient(config Config, deps Deps) APIClient {

	router := chi.NewRouter()
	middleware.Default(router)

	router.Route("/hello-world", func(router chi.Router) {
		router.Get("/", route.HelloWorld())
	})

	return APIClient{
		config: config,
		deps:   deps,
		router: router,
	}
}

func (api APIClient) ListenAndServe(ctx context.Context) error {

	log.Info(ctx, "Listening", log.FmtString(api.config.Port, "port"))
	if err := http.ListenAndServe(api.config.Port, api.router); err != nil {
		return errors.Wrap(err, "Failed listening and serving")
	}
	return nil
}
