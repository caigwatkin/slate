/*
Copyright 2018 Cai Gwatkin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"context"
	"net/http"
	pkg_http "slate/internal/pkg/http"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secrets"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type Client struct {
	config        Config
	httpClient    pkg_http.Client
	logClient     log.Client
	secretsClient *secrets.Client
	router        *chi.Mux
	serviceName   string
}

type Config struct {
	Env          string
	GCPProjectID string
	Port         string
}

func NewClient(config Config, httpClient pkg_http.Client, logClient log.Client, secretsClient *secrets.Client) Client {
	apiClient := Client{
		config:        config,
		httpClient:    httpClient,
		logClient:     logClient,
		secretsClient: secretsClient,
		router:        chi.NewRouter(),
		serviceName:   "slate",
	}
	pathForHealthEndpoint := "/health"
	apiClient.loadMiddleware(pathForHealthEndpoint)
	apiClient.loadEndpoints(pathForHealthEndpoint)
	return apiClient
}

func (api Client) loadMiddleware(pathForHealthEndpoint string) {
	api.httpClient.MiddlewareDefaults(api.router, []string{pathForHealthEndpoint})
}

func (api Client) ListenAndServe(ctx context.Context) error {
	api.logClient.Info(ctx, "Listening and serving", log.FmtString(api.config.Port, "port"))

	if err := http.ListenAndServe(api.config.Port, api.router); err != nil {
		return errors.Wrap(err, "Failed listening and serving")
	}
	return nil
}
