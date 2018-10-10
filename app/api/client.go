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
	"net/http"

	go_http "github.com/caigwatkin/go/http"
	"github.com/caigwatkin/slate/app/routes"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type Client interface {
	ListenAndServe() error
}

type client struct {
	config       Config
	httpClient   go_http.Client
	routesClient routes.Client
	router       *chi.Mux
}

type Config struct {
	Env          string
	GCPProjectID string
	Port         string
}

func NewClient(config Config, httpClient go_http.Client, routesClient routes.Client) Client {
	c := client{
		config:       config,
		httpClient:   httpClient,
		routesClient: routesClient,
		router:       chi.NewRouter(),
	}
	pathForHealthEndpoint := "/health"
	c.loadMiddleware(pathForHealthEndpoint)
	c.loadEndpoints(pathForHealthEndpoint)
	return c
}

func (c client) loadMiddleware(pathForHealthEndpoint string) {
	c.httpClient.MiddlewareDefaults(c.router, []string{pathForHealthEndpoint})
}

func (c client) ListenAndServe() error {
	if err := http.ListenAndServe(c.config.Port, c.router); err != nil {
		return errors.Wrap(err, "Failed listening and serving")
	}
	return nil
}
