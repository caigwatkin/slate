package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	go_errors "github.com/caigwatkin/go/errors"
	go_secrets "github.com/caigwatkin/go/secrets"
	"google.golang.org/api/option"
)

type Client interface {
	Close()
}

type client struct {
	firestoreClient *firestore.Client
}

func NewClient(ctx context.Context, secretsClient go_secrets.Client) (Client, error) {
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
