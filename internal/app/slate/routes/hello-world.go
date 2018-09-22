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
	pkg_http "github.com/caigwatkin/slate/internal/pkg/http"
	"github.com/caigwatkin/slate/internal/pkg/log"
	"net/http"
)

func ReadHelloWorld(httpClient pkg_http.Client, logClient log.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logClient.Info(ctx, "Reading")

		b, err := json.MarshalIndent(map[string]string{
			"hello": "world",
		}, "", "\t")
		if err != nil {
			httpClient.RenderErrorOrStatus(ctx, w, err)
		}
		httpClient.RenderContentJSON(ctx, w, b)
	}
}
