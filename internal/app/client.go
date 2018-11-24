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

package app

import (
	"context"

	go_log "github.com/caigwatkin/go/log"
	"github.com/caigwatkin/slate/internal/lib/dto"
	"github.com/caigwatkin/slate/internal/lib/firestore"
)

type Client interface {
	CreateGreeting(ctx context.Context, d dto.CreateGreeting) (string, error)
	DeleteGreeting(ctx context.Context, d dto.DeleteGreeting) error
	ReadGreeting(ctx context.Context, d dto.ReadGreeting) ([]byte, error)
}

type client struct {
	firestoreClient firestore.Client
	logClient       go_log.Client
}

func NewClient(firestoreClient firestore.Client, logClient go_log.Client) Client {
	return client{
		firestoreClient: firestoreClient,
		logClient:       logClient,
	}
}
