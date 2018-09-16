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
	"slate/internal/pkg/log"
)

func logErrorMarshallingJSONBody(ctx context.Context, code int, headers map[string]string) {
	log.Error(ctx, "Failed marshalling JSON for response body",
		log.FmtInt(code, "status code"),
		log.FmtString(http.StatusText(code), "status text"),
		log.FmtAny(headers, "headers"),
	)
}

func logErrorWritingBody(ctx context.Context, code int, headers map[string]string, body []byte) {
	log.Error(ctx, "Failed writing body to response",
		log.FmtInt(code, "status code"),
		log.FmtString(http.StatusText(code), "status text"),
		log.FmtAny(headers, "headers"),
		log.FmtBytes(body, "body"),
	)
}

func logInfoResponse(ctx context.Context, code int, headers map[string]string, lenBody int, body []byte) {
	log.Info(ctx, "HTTP Response",
		log.FmtInt(code, "status code"),
		log.FmtString(http.StatusText(code), "status text"),
		log.FmtAny(headers, "headers"),
		log.FmtInt(lenBody, "lenBody"),
		log.FmtBytes(body, "body"),
	)
}
