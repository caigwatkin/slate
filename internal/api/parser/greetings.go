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

package parser

import (
	"net/http"

	go_errors "github.com/caigwatkin/go/errors"
	go_log "github.com/caigwatkin/go/log"
	"github.com/caigwatkin/slate/internal/lib/dto"
	"github.com/go-chi/chi"
)

func (c client) CreateGreeting(r *http.Request) dto.CreateGreeting {
	ctx := r.Context()
	c.logClient.Info(ctx, "Parsing")
	d := dto.CreateGreeting{
		Message: "hello world",
	}
	c.logClient.Info(ctx, "Parsed", go_log.FmtAny(d, "d"))
	return d
}

func (c client) ReadGreeting(r *http.Request) (*dto.ReadGreeting, error) {
	ctx := r.Context()
	c.logClient.Info(ctx, "Parsing")
	id := chi.URLParam(r, "greeting_id")
	if id == "" {
		return nil, go_errors.NewStatus(http.StatusBadRequest, "greeting_id")
	}
	d := dto.ReadGreeting{
		ID: id,
	}
	c.logClient.Info(ctx, "Parsed", go_log.FmtAny(d, "d"))
	return &d, nil
}

func (c client) DeleteGreeting(r *http.Request) (*dto.DeleteGreeting, error) {
	ctx := r.Context()
	c.logClient.Info(ctx, "Parsing")
	id := chi.URLParam(r, "greeting_id")
	if id == "" {
		return nil, go_errors.NewStatus(http.StatusBadRequest, "greeting_id")
	}
	d := dto.DeleteGreeting{
		ID: id,
	}
	c.logClient.Info(ctx, "Parsed", go_log.FmtAny(d, "d"))
	return &d, nil
}
