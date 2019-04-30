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
)

type Client interface {
	ParseCreateGreeting(r *http.Request) dto.CreateGreeting
	ParseReadGreeting(r *http.Request) (*dto.ReadGreeting, error)
	ParseDeleteGreeting(r *http.Request) (*dto.DeleteGreeting, error)
}

type client struct {
	logClient go_log.Client
}

func NewClient(logClient go_log.Client) Client {
	return client{
		logClient: logClient,
	}
}

func missingURLParam(p string) error {
	return go_errors.Statusf(http.StatusBadRequest, "Missing URL param %q", p)
}
