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

package router

import (
	"net/http"

	go_errors "github.com/caigwatkin/go/errors"
)

func (c client) CreateGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c.logClient.Info(ctx, "Creating")
		location, err := c.appClient.CreateGreeting(ctx, c.parserClient.CreateGreeting(r))
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed creating greeting using app client"))
			return
		}
		c.httpClient.RenderCreated(ctx, w, location)
	}
}

func (c client) ReadGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c.logClient.Info(ctx, "Reading")
		d, err := c.parserClient.ReadGreeting(r)
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed parsing read greeting using parser client"))
			return
		}
		greeting, err := c.appClient.ReadGreeting(ctx, *d)
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed reading greeting using app client"))
			return
		}
		c.httpClient.RenderContentJSON(ctx, w, greeting)
	}
}

func (c client) DeleteGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c.logClient.Info(ctx, "Deleting")
		d, err := c.parserClient.DeleteGreeting(r)
		if err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed parsing delete greeting using parser client"))
			return
		}
		if err := c.appClient.DeleteGreeting(ctx, *d); err != nil {
			c.httpClient.RenderErrorOrStatus(ctx, w, go_errors.Wrap(err, "Failed deleting greeting using app client"))
			return
		}
		c.httpClient.RenderNoContent(ctx, w)
	}
}
