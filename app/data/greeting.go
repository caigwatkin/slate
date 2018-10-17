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

package data

import (
	"context"
	"encoding/json"
	"net/http"

	go_errors "github.com/caigwatkin/go/errors"
	go_log "github.com/caigwatkin/go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Greeting struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func (g Greeting) Render() ([]byte, error) {
	b, err := json.MarshalIndent(g, "", "\t")
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed unmarshalling data bytes into Greeting")
	}
	return b, nil
}

func (c client) CreateGreeting(ctx context.Context, message string) (id string, err error) {
	c.logClient.Info(ctx, "Creating", go_log.FmtString(message, "message"))
	g, err := createGreetingData(message)
	if err != nil {
		err = go_errors.Wrap(err, "Failed converting greeting to map")
		return
	}
	documentRef, _, err := c.firestoreClient.Collection("greeting").Add(ctx, g)
	if err != nil {
		err = go_errors.Wrap(err, "Failed adding new greeting to collection")
		return
	}
	id = documentRef.ID
	c.logClient.Info(ctx, "Created", go_log.FmtString(id, "id"))
	return
}

func createGreetingData(message string) (map[string]interface{}, error) {
	g := Greeting{
		Message: message,
	}
	b, err := json.Marshal(g)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed unmarshalling greeting")
	}
	gm := make(map[string]interface{})
	if err := json.Unmarshal(b, &gm); err != nil {
		return nil, go_errors.Wrap(err, "Failed unmarshalling greeting bytes into map[string]interface{}")
	}
	return gm, nil
}

func (c client) DeleteGreeting(ctx context.Context, id string) error {
	c.logClient.Info(ctx, "Deleting", go_log.FmtString(id, "id"))
	documentRef := c.firestoreClient.Collection("greeting").Doc(id)
	if documentRef == nil {
		return go_errors.Errorf("Failed getting document reference for id %q", id)
	}
	if _, err := documentRef.Delete(ctx); err != nil {
		return go_errors.Errorf("Failed deleting document by reference")
	}
	c.logClient.Info(ctx, "Deleted")
	return nil
}

func (c client) ReadGreeting(ctx context.Context, id string) (*Greeting, error) {
	c.logClient.Info(ctx, "Reading", go_log.FmtString(id, "id"))
	documentRef := c.firestoreClient.Collection("greeting").Doc(id)
	if documentRef == nil {
		return nil, go_errors.Errorf("Failed getting document reference for id %q", id)
	}
	documentSnapshot, err := documentRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, go_errors.NewStatus(http.StatusNotFound, "greeting does not exist")
		}
		return nil, go_errors.Wrap(err, "Failed getting snapshot of document")
	}
	greeting, err := greetingFromDocSnapshotData(documentSnapshot.Data())
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed generating Greeting from document snapshot data")
	}
	c.logClient.Info(ctx, "Read", go_log.FmtAny(greeting, "greeting"))
	return greeting, nil
}

func greetingFromDocSnapshotData(data map[string]interface{}) (*Greeting, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed unmarshalling data")
	}
	var g Greeting
	if err := json.Unmarshal(b, &g); err != nil {
		return nil, go_errors.Wrap(err, "Failed unmarshalling data bytes into Greeting")
	}
	return &g, nil
}
