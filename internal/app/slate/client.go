package slate

import (
	"context"
	"net/http"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secret"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type RESTAPIClient struct {
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

func NewRESTAPIClient(config Config, deps Deps) RESTAPIClient {

	router := chi.NewRouter()
	apiClient := RESTAPIClient{
		config: config,
		deps:   deps,
		router: router,
	}
	apiClient.middleware()
	apiClient.api()

	return apiClient
}

func (api RESTAPIClient) ListenAndServe(ctx context.Context) error {

	log.Info(ctx, "Listening and serving", log.FmtString(api.config.Port, "port"))
	if err := http.ListenAndServe(api.config.Port, api.router); err != nil {
		return errors.Wrap(err, "Failed listening and serving")
	}
	return nil
}
