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
	"net/http"
	"slate/internal/pkg/context"
	"strings"

	"github.com/google/uuid"
)

const (
	HeaderValXSlateTest = "5c1bca85-9e09-4af4-96ac-7f353265838c"
)

// PopulateContext adds custom ctx params
func PopulateContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithCorrelationID(r.Context(), uuid.New().String())
		if v, ok := r.Header["X-Slate-Correlation-Id"]; ok {
			ctx = context.WithCorrelationIDAppend(ctx, strings.Join(v, ","))
		}
		var t bool
		if v, ok := r.Header["X-Slate-Test"]; ok {
			t = strings.Join(v, ",") == HeaderValXSlateTest
		}
		ctx = context.WithTest(ctx, t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
