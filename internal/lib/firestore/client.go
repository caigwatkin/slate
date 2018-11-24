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

package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	go_errors "github.com/caigwatkin/go/errors"
	go_log "github.com/caigwatkin/go/log"
	go_secrets "github.com/caigwatkin/go/secrets"
	"github.com/caigwatkin/slate/internal/lib/dto"
	"google.golang.org/api/option"
)

type Client interface {
	Close()

	CreateGreeting(ctx context.Context, d dto.CreateGreeting) (location string, err error)
	ReadGreeting(ctx context.Context, d dto.ReadGreeting) (*Greeting, error)
	DeleteGreeting(ctx context.Context, d dto.DeleteGreeting) error
	RenderGreeting(g Greeting) ([]byte, error)
}

type client struct {
	firestoreClient *firestore.Client
	logClient       go_log.Client
}

func NewClient(ctx context.Context, logClient go_log.Client, secretsClient go_secrets.Client) (Client, error) {
	s, err := secretsClient.Secret(secretDomainSlateAPI, secretTypeServiceAccountKey)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed retrieving GCP service account key from secrets")
	}
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(s))
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed creating firebase app")
	}
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed creating firestore client")
	}
	return client{
		firestoreClient: firestoreClient,
		logClient:       logClient,
	}, nil
}

const (
	secretDomainSlateAPI        = "slate_api_dev"
	secretTypeServiceAccountKey = "service_account_key"
)

func RequiredSecrets() go_secrets.Required {
	return map[string][]string{
		secretDomainSlateAPI: []string{
			secretTypeServiceAccountKey,
		},
	}
}

func (c client) Close() {
	c.firestoreClient.Close()
}
