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
	"slate/internal/pkg/log"
	"strings"

	"github.com/google/uuid"
)

const (
	HeaderKeyXSlateCorrelationID = "X-Slate-Correlation-Id"
	HeaderKeyXSlateTest          = "X-Slate-Test"

	HeaderValXSlateTest = "5c1bca85-9e09-4af4-96ac-7f353265838c"
)

func PopulateContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pkg_context.WithCorrelationID(r.Context(), uuid.New().String())
		if v, ok := r.Header[HeaderKeyXSlateCorrelationID]; ok {
			ctx = pkg_context.WithCorrelationIDAppend(ctx, strings.Join(v, ","))
		}
		var t bool
		if v, ok := r.Header[HeaderKeyXSlateTest]; ok {
			t = strings.Join(v, ",") == HeaderValXSlateTest
		}
		ctx = pkg_context.WithTest(ctx, t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type RequestLogger struct {
	exclude []string
}

func NewRequestLogger(exclude []string) *RequestLogger {
	return &RequestLogger{
		exclude: exclude,
	}
}

func (l *RequestLogger) Info(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		var exclude bool
		for _, v := range l.exclude {
			if url == v {
				exclude = true
				break
			}
		}
		if !exclude {
			log.Info(r.Context(), "HTTP Request Received",
				log.FmtString(r.URL.String(), "URL"),
				log.FmtString(r.Method, "Method"),
				log.FmtAny(r.Header, "Header"),
			)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
