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
	"encoding/json"
	"net/http"

	go_errors "github.com/caigwatkin/go/errors"
	go_log "github.com/caigwatkin/go/log"
	"github.com/caigwatkin/slate/internal/lib/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Greeting struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func newGreeting(d dto.CreateGreeting) Greeting {
	return Greeting{
		Message: d.UserInput.Message,
	}
}

func (c client) RenderGreeting(g Greeting) ([]byte, error) {
	b, err := json.Marshal(g)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed marshalling greeting into bytes")
	}
	return b, nil
}

func (g Greeting) toMap() (map[string]interface{}, error) {
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

func (c client) CreateGreeting(ctx context.Context, d dto.CreateGreeting) (location string, err error) {
	c.logClient.Info(ctx, "Creating", go_log.FmtAny(d, "d"))
	g, err := newGreeting(d).toMap()
	if err != nil {
		err = go_errors.Wrap(err, "Failed converting greeting to map")
		return
	}
	documentRef, _, err := c.firestoreClient.Collection("greeting").Add(ctx, g)
	if err != nil {
		err = go_errors.Wrap(err, "Failed adding new greeting to collection")
		return
	}
	location = documentRef.ID
	c.logClient.Info(ctx, "Created", go_log.FmtString(location, "location"))
	return
}

func (c client) ReadGreeting(ctx context.Context, d dto.ReadGreeting) (*Greeting, error) {
	c.logClient.Info(ctx, "Reading", go_log.FmtAny(d, "d"))
	documentRef := c.firestoreClient.Collection("greeting").Doc(d.ID)
	if documentRef == nil {
		return nil, go_errors.Errorf("Failed getting document reference for d.ID %q", d.ID)
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

func (c client) DeleteGreeting(ctx context.Context, d dto.DeleteGreeting) error {
	c.logClient.Info(ctx, "Deleting", go_log.FmtAny(d, "d"))
	documentRef := c.firestoreClient.Collection("greeting").Doc(d.ID)
	if documentRef == nil {
		return go_errors.Errorf("Failed getting document reference for id %q", d.ID)
	}
	if _, err := documentRef.Delete(ctx); err != nil {
		return go_errors.Errorf("Failed deleting document by reference")
	}
	c.logClient.Info(ctx, "Deleted")
	return nil
}
