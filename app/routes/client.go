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

package routes

import (
	"net/http"

	go_http "github.com/caigwatkin/go/http"
	go_log "github.com/caigwatkin/go/log"
	"github.com/caigwatkin/slate/app/data"
)

type Client interface {
	HealthClient
	GreetingsClient
}

type GreetingsClient interface {
	CreateGreeting() http.HandlerFunc
	ReadGreeting() http.HandlerFunc
}

type HealthClient interface {
	Health() http.HandlerFunc
}

type client struct {
	dataClient  data.Client
	httpClient  go_http.Client
	logClient   go_log.Client
	serviceName string
}

func NewClient(dataClient data.Client, httpClient go_http.Client, logClient go_log.Client, serviceName string) Client {
	return client{
		dataClient:  dataClient,
		httpClient:  httpClient,
		logClient:   logClient,
		serviceName: serviceName,
	}
}
