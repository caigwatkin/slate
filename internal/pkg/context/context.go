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

package context

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

const (
	CorrelationIDBackground = "BACKGROUND"
	CorrelationIDStartUp    = "START_UP"
)

// Background new ctx with background correlation ID
func Background() context.Context {
	ctx := WithCorrelationID(context.Background(), CorrelationIDBackground)
	return WithTest(ctx, false)
}

// StartUp new ctx with start up correlation ID
func StartUp() context.Context {
	ctx := WithCorrelationID(context.Background(), CorrelationIDStartUp)
	return WithTest(ctx, false)
}

// New ctx with values from existing ctx
// -> useful for goroutines that should continue even after the parent context has ended
func New(ctx context.Context) context.Context {
	c := WithCorrelationID(context.Background(), uuid.New().String())
	c = WithCorrelationIDAppend(c, CorrelationID(ctx))
	c = WithTest(c, Test(ctx))
	return WithTest(c, false)
}

type key int

const (
	keyCorrelationID key = iota
	keyTest          key = iota
)

// CorrelationID from ctx
func CorrelationID(ctx context.Context) string {
	if v, ok := ctx.Value(keyCorrelationID).(string); ok {
		return v
	}
	return ""
}

// WithCorrelationID in new ctx
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, keyCorrelationID, correlationID)
}

// WithCorrelationIDAppend in new ctx
func WithCorrelationIDAppend(ctx context.Context, correlationID string) context.Context {
	if cID := CorrelationID(ctx); cID != CorrelationIDBackground {
		correlationID = fmt.Sprintf("%s,%s", cID, correlationID)
	}
	return WithCorrelationID(ctx, correlationID)
}

// Test from ctx
func Test(ctx context.Context) bool {
	if v, ok := ctx.Value(keyTest).(bool); ok {
		return v
	}
	return false
}

// WithTest in new ctx
func WithTest(ctx context.Context, test bool) context.Context {
	return context.WithValue(ctx, keyTest, test)
}
