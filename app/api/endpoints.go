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
	"github.com/go-chi/chi"
)

func (c *client) loadEndpoints(pathForHealthEndpoint string) {
	router := c.router
	router.Get(pathForHealthEndpoint, c.routesClient.Health())
	router.Route("/greetings", func(router chi.Router) {
		router.Post("/", c.routesClient.CreateGreeting())
		router.Route("/{greeting_id}", func(router chi.Router) {
			router.Get("/", c.routesClient.ReadGreeting())
			router.Delete("/", c.routesClient.DeleteGreeting())
		})
	})
}
