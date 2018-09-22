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
	"slate/internal/app/slate/routes"

	"github.com/go-chi/chi"
)

func (api *Client) loadEndpoints(pathForHealthEndpoint string) {
	api.router.Get(pathForHealthEndpoint, routes.Health(api.httpClient, api.serviceName))
	api.router.Route("/hello-world", func(router chi.Router) {
		router.Get("/", routes.ReadHelloWorld(api.httpClient, api.logClient))
	})
}
