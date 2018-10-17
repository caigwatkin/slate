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

package routes

import (
	"net/http"

	go_errors "github.com/caigwatkin/go/errors"
	"github.com/go-chi/chi"
)

func (c client) CreateGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c.logClient.Info(ctx, "Creating")
		location, err := c.dataClient.CreateGreeting(ctx, "hello world")
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed creating greeting using data client"))
			return
		}
		c.httpClient.RenderCreated(ctx, w, location)
	}
}

func (c client) DeleteGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c.logClient.Info(ctx, "Deleting")
		if err := c.dataClient.DeleteGreeting(ctx, chi.URLParam(r, "greeting_id")); err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed deleting greeting using data client"))
			return
		}
		c.httpClient.RenderNoContent(ctx, w)
	}
}

func (c client) ReadGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c.logClient.Info(ctx, "Reading")
		greeting, err := c.dataClient.ReadGreeting(ctx, chi.URLParam(r, "greeting_id"))
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed reading greeting from data client"))
			return
		}
		b, err := greeting.Render()
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed rendering greeting"))
			return
		}
		c.httpClient.RenderContentJSON(ctx, w, b)
	}
}
