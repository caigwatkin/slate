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
	"slate/internal/pkg/http/render"
	"slate/internal/pkg/log"
)

func Ping() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Info(ctx, "Pinging")

		b, err := json.MarshalIndent(map[string]interface{}{
			"hello": "world",
		}, "", "\t")
		if err != nil {
			render.ErrorOrStatus(ctx, w, err)
		}
		render.ContentJSON(ctx, w, b)
	}
}
