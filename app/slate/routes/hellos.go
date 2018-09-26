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
	"encoding/json"
	"net/http"

	go_http "github.com/caigwatkin/go/http"
	go_log "github.com/caigwatkin/go/log"
	"github.com/go-chi/chi"
)

func CreateHello(httpClient go_http.Client, logClient go_log.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logClient.Info(ctx, "Creating")

		httpClient.RenderCreated(ctx, w, "some_id")
	}
}

func ReadHelloByID(httpClient go_http.Client, logClient go_log.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logClient.Info(ctx, "Reading")

		helloID := chi.URLParam(r, "hello_id")
		b, err := json.MarshalIndent(map[string]string{
			"hello": "world",
			"id":    helloID,
		}, "", "\t")
		if err != nil {
			httpClient.RenderErrorOrStatus(ctx, w, err)
		}
		httpClient.RenderContentJSON(ctx, w, b)
	}
}
