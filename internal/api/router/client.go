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

package router

import (
	"net/http"

	go_http "github.com/caigwatkin/go/http"
	go_log "github.com/caigwatkin/go/log"
	"github.com/caigwatkin/slate/internal/api/parser"
	"github.com/caigwatkin/slate/internal/app"
)

type Client interface {
	Health() http.HandlerFunc

	CreateGreeting() http.HandlerFunc
	ReadGreeting() http.HandlerFunc
	DeleteGreeting() http.HandlerFunc
}

type client struct {
	config       Config
	appClient    app.Client
	httpClient   go_http.Client
	logClient    go_log.Client
	parserClient parser.Client
}

type Config struct {
	ServiceName string
}

func NewClient(config Config, appClient app.Client, httpClient go_http.Client, logClient go_log.Client, parserClient parser.Client) Client {
	return client{
		config:       config,
		appClient:    appClient,
		httpClient:   httpClient,
		logClient:    logClient,
		parserClient: parserClient,
	}
}
