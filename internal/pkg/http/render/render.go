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
	"encoding/json"
	"net/http"
	"slate/internal/pkg/http/headers"
	pkg_httpstatus "slate/internal/pkg/http/status"
)

func ContentJSON(ctx context.Context, headersClient headers.Client, w http.ResponseWriter, body []byte) {
	headers := setHeadersInclDefaults(ctx, headersClient, w, map[string]string{
		"Content-Type": "application/json",
	})
	body = append(body, byte('\n'))
	code := http.StatusOK
	w.WriteHeader(code)
	lenBody, err := w.Write(body)
	if err != nil {
		logErrorWritingBody(ctx, code, headers, body)
		return
	}
	logInfoResponse(ctx, code, headers, lenBody, nil)
}

func Health(ctx context.Context, headersClient headers.Client, w http.ResponseWriter, serviceName string) {
	headers := setHeadersInclDefaults(ctx, headersClient, w, map[string]string{
		"Content-Type": "application/json",
	})
	body, err := json.MarshalIndent(map[string]string{
		"service": serviceName,
		"status":  "OK",
	}, "", "\t")
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		logErrorMarshallingJSONBody(ctx, code, headers)
		return
	}
	body = append(body, byte('\n'))
	code := http.StatusOK
	w.WriteHeader(code)
	if _, err := w.Write(body); err != nil {
		logErrorWritingBody(ctx, code, headers, body)
		return
	}
}

func Status(ctx context.Context, headersClient headers.Client, w http.ResponseWriter, s pkg_httpstatus.Status) {
	h := setHeadersInclDefaults(ctx, headersClient, w, map[string]string{
		"Content-Type": "application/json",
	})
	body := s.Render()
	w.WriteHeader(s.Code)
	lenBody, err := w.Write(body)
	if err != nil {
		logErrorWritingBody(ctx, s.Code, h, body)
		return
	}
	logInfoResponse(ctx, s.Code, h, lenBody, body)
}

func ErrorOrStatus(ctx context.Context, headersClient headers.Client, w http.ResponseWriter, err error) {
	if v, ok := err.(pkg_httpstatus.Status); ok {
		Status(ctx, headersClient, w, v)
		return
	}
	h := setHeadersInclDefaults(ctx, headersClient, w, nil)
	code := http.StatusInternalServerError
	w.WriteHeader(code)
	logInfoResponse(ctx, code, h, 0, nil)
}
