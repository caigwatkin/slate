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
	"slate/internal/pkg/http/headers"
	"slate/internal/pkg/log"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

func Defaults(r *chi.Mux, headersClient headers.Client, excludePathsForLogInfoRequests []string) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 30))
	r.Use(middleware.URLFormat)
	r.Use(PopulateContext(headersClient))
	r.Use(LogInfoRequests(excludePathsForLogInfoRequests))
	r.Use(middleware.DefaultCompress)
}

func PopulateContext(headersClient headers.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := pkg_context.WithCorrelationID(r.Context(), uuid.New().String())
			if v, ok := r.Header[headersClient.CorrelationIDKey()]; ok {
				ctx = pkg_context.WithCorrelationIDAppend(ctx, strings.Join(v, ","))
			}
			var t bool
			if v, ok := r.Header[headersClient.TestKey()]; ok {
				t = strings.Join(v, ",") == headers.TestValDefault
			}
			ctx = pkg_context.WithTest(ctx, t)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func LogInfoRequests(excludePaths []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.String()
			var exclude bool
			for _, v := range excludePaths {
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
