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

package middleware

import (
	"net/http"
	pkg_context "slate/internal/pkg/context"
	"slate/internal/pkg/http/constants"
	"slate/internal/pkg/log"
	"strings"

	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

func Default(r *chi.Mux, requestLoggingExclusionPaths []string) {
	r.Use(chi_middleware.RequestID)
	r.Use(chi_middleware.DefaultCompress)
	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.Recoverer)
	r.Use(chi_middleware.URLFormat)
	r.Use(populateContext)
	r.Use(infoLogRequests(requestLoggingExclusionPaths))
}

func populateContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pkg_context.WithCorrelationID(r.Context(), uuid.New().String())
		if v, ok := r.Header[constants.HeaderKeyXSlateCorrelationID]; ok {
			ctx = pkg_context.WithCorrelationIDAppend(ctx, strings.Join(v, ","))
		}
		var t bool
		if v, ok := r.Header[constants.HeaderKeyXSlateTest]; ok {
			t = strings.Join(v, ",") == constants.HeaderValXSlateTest
		}
		ctx = pkg_context.WithTest(ctx, t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func infoLogRequests(exclusionPaths []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.String()
			var exclude bool
			for _, v := range exclusionPaths {
				if url == v {
					exclude = true
					break
				}
			}
			if !exclude {
				log.Info(r.Context(), "HTTP Request",
					log.FmtString(r.URL.String(), "URL"),
					log.FmtString(r.Method, "Method"),
					log.FmtAny(r.Header, "Header"),
				)
			}
			next.ServeHTTP(w, r)
		})
	}
}
