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

package render

import (
	"context"
	"net/http"
	pkg_context "slate/internal/pkg/context"
	"slate/internal/pkg/http/constants"
	http_status "slate/internal/pkg/http/status"
	"slate/internal/pkg/log"
)

func ContentJSON(ctx context.Context, w http.ResponseWriter, body []byte) {
	body = append(body, byte('\n'))
	h := defaultHeaders(ctx)
	h["Content-Type"] = "application/json"
	for k, v := range h {
		w.Header().Set(k, v)
	}
	lenBody, err := w.Write(body)
	if err != nil {
		log.Error(ctx, "Failed writing body to response",
			log.FmtAny(h, "headers"),
			log.FmtBytes(body, "body"),
		)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Info(ctx, "HTTP Response",
		log.FmtInt(http.StatusOK, "status code"),
		log.FmtString(http.StatusText(http.StatusOK), "status text"),
		log.FmtAny(h, "headers"),
		log.FmtInt(lenBody, "lenBody"),
		log.FmtBytes(body, "body"),
	)
}

func ErrorOrStatus(ctx context.Context, w http.ResponseWriter, err error) {
	if v, ok := err.(http_status.Status); ok {
		Status(ctx, w, v)
	}
	h := defaultHeaders(ctx)
	w.WriteHeader(http.StatusInternalServerError)
	log.Info(ctx, "HTTP Response",
		log.FmtInt(http.StatusInternalServerError, "status code"),
		log.FmtString(http.StatusText(http.StatusInternalServerError), "status text"),
		log.FmtAny(h, "headers"),
	)
}

func Status(ctx context.Context, w http.ResponseWriter, s http_status.Status) {
	h := defaultHeaders(ctx)
	h["Content-Type"] = "application/json"
	for k, v := range h {
		w.Header().Set(k, v)
	}
	w.WriteHeader(s.Code)
	body := s.Render()
	lenBody, err := w.Write(body)
	if err != nil {
		log.Error(ctx, "Failed writing body to response",
			log.FmtAny(h, "headers"),
			log.FmtBytes(body, "body"),
		)
	}
	log.Info(ctx, "HTTP Response",
		log.FmtInt(s.Code, "status code"),
		log.FmtString(http.StatusText(s.Code), "status text"),
		log.FmtAny(h, "headers"),
		log.FmtInt(lenBody, "lenBody"),
		log.FmtBytes(body, "body"),
	)
}

func defaultHeaders(ctx context.Context) map[string]string {
	headers := map[string]string{
		constants.HeaderKeyXSlateCorrelationID: pkg_context.CorrelationID(ctx),
	}
	if pkg_context.Test(ctx) {
		headers[constants.HeaderKeyXSlateTest] = constants.HeaderValXSlateTest
	}
	return headers
}
