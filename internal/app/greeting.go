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

	go_errors "github.com/caigwatkin/go/errors"
	go_log "github.com/caigwatkin/go/log"
	"github.com/caigwatkin/slate/internal/lib/dto"
)

func (c client) CreateGreeting(ctx context.Context, d dto.CreateGreeting) (string, error) {
	c.logClient.Info(ctx, "Creating", go_log.FmtAny(d, "d"))
	location, err := c.firestoreClient.CreateGreeting(ctx, d)
	if err != nil {
		return "", go_errors.Wrap(err, "Failed creating greeting using firestore client")
	}
	c.logClient.Info(ctx, "Created", go_log.FmtString(location, "location"))
	return location, nil
}

func (c client) DeleteGreeting(ctx context.Context, d dto.DeleteGreeting) error {
	c.logClient.Info(ctx, "Deleting", go_log.FmtAny(d, "d"))
	if err := c.firestoreClient.DeleteGreeting(ctx, d); err != nil {
		return go_errors.Wrap(err, "Failed deleting greeting using firestore client")
	}
	c.logClient.Info(ctx, "Deleted")
	return nil
}

func (c client) ReadGreeting(ctx context.Context, d dto.ReadGreeting) ([]byte, error) {
	c.logClient.Info(ctx, "Reading", go_log.FmtAny(d, "d"))
	g, err := c.firestoreClient.ReadGreeting(ctx, d)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed reading greeting from firestore client")
	}
	b, err := c.firestoreClient.RenderGreeting(*g)
	if err != nil {
		return nil, go_errors.Wrap(err, "Failed rendering greeting")
	}
	c.logClient.Info(ctx, "Read", go_log.FmtBytes(b, "b"))
	return b, nil
}
