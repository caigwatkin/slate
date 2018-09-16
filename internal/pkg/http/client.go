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

package http

import (
	"context"
	"net/http"
	"slate/internal/pkg/http/headers"
	"slate/internal/pkg/http/middleware"
	"slate/internal/pkg/http/render"

	"github.com/go-chi/chi"
)

type Client interface {
	RenderContentJSON(ctx context.Context, w http.ResponseWriter, body []byte)
	RenderHealth(ctx context.Context, w http.ResponseWriter, serviceName string)
	RenderErrorOrStatus(ctx context.Context, w http.ResponseWriter, err error)
	MiddlewareDefaults(r *chi.Mux, excludePathsForLogInfoRequests []string)
}

type client struct {
	headersClient headers.Client
}

// NewClient for http
//
// Service name should be in canonical case
// Use an empty string to use default keys
func NewClient(serviceName string) Client {
	return client{
		headersClient: headers.NewClient(serviceName),
	}
}

// RenderContentJSON in response
func (c client) RenderContentJSON(ctx context.Context, w http.ResponseWriter, body []byte) {
	render.ContentJSON(ctx, c.headersClient, w, body)
}

// RenderHealth in response
func (c client) RenderHealth(ctx context.Context, w http.ResponseWriter, serviceName string) {
	render.Health(ctx, c.headersClient, w, serviceName)
}

// RenderErrorOrStatus in response
func (c client) RenderErrorOrStatus(ctx context.Context, w http.ResponseWriter, err error) {
	render.ErrorOrStatus(ctx, c.headersClient, w, err)
}

// MiddlewareDefaults for request handling
func (c client) MiddlewareDefaults(r *chi.Mux, excludePathsForLogInfoRequests []string) {
	middleware.Defaults(r, c.headersClient, excludePathsForLogInfoRequests)
}
