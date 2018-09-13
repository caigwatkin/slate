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

func Background() context.Context {
	ctx := WithCorrelationID(context.Background(), CorrelationIDBackground)
	return WithTest(ctx, false)
}

func StartUp() context.Context {
	ctx := WithCorrelationID(context.Background(), CorrelationIDStartUp)
	return WithTest(ctx, false)
}

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

func CorrelationID(ctx context.Context) string {
	if v, ok := ctx.Value(keyCorrelationID).(string); ok {
		return v
	}
	return ""
}

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, keyCorrelationID, correlationID)
}

func WithCorrelationIDAppend(ctx context.Context, correlationID string) context.Context {
	if cID := CorrelationID(ctx); cID != CorrelationIDBackground {
		correlationID = fmt.Sprintf("%s,%s", cID, correlationID)
	}
	return WithCorrelationID(ctx, correlationID)
}

func Test(ctx context.Context) bool {
	if v, ok := ctx.Value(keyTest).(bool); ok {
		return v
	}
	return false
}

func WithTest(ctx context.Context, test bool) context.Context {
	return context.WithValue(ctx, keyTest, test)
}
