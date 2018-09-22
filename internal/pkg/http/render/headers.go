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
	pkg_context "github.com/caigwatkin/slate/internal/pkg/context"
	pkg_headers "github.com/caigwatkin/slate/internal/pkg/http/headers"
	"net/http"
)

func setHeadersInclDefaults(ctx context.Context, headersClient pkg_headers.Client, w http.ResponseWriter, headers map[string]string) map[string]string {
	h := map[string]string{
		headersClient.CorrelationIDKey(): pkg_context.CorrelationID(ctx),
	}
	if pkg_context.Test(ctx) {
		h[headersClient.TestKey()] = pkg_headers.TestValDefault
	}
	for k, v := range headers {
		h[k] = v
	}
	for k, v := range h {
		w.Header().Set(k, v)
	}
	return h
}
