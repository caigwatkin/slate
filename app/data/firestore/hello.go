package firestore

import (
	"context"
	"time"

	"github.com/caigwatkin/go/errors"
)

func (c client) CreateHello(ctx context.Context) (string, error) {
	documentRef, _, err := c.firestoreClient.Collection("hello").Add(ctx, map[string]interface{}{
		"date_created": time.Now().UTC(),
		"message":      "world",
	})
	if err != nil {
		return "", errors.Wrap(err, "Failed adding new hello to collection")
	}
	return documentRef.ID, nil
}

func (c client) ReadHello(ctx context.Context, id string) (map[string]interface{}, error) {
	documentRef := c.firestoreClient.Collection("hello").Doc(id)
	if documentRef == nil {
		return nil, errors.Errorf("Failed getting document reference for id %q", id)
	}
	documentSnapshot, err := documentRef.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed getting snapshot of document")
	}
	return documentSnapshot.Data(), nil
}
