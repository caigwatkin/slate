package log

import (
	"context"
)

var (
	MockClientSuccess Client = mockClientSuccess{}
)

type mockClientSuccess struct{}

func (c mockClientSuccess) Debug(_ context.Context, _ string, _ ...Field) {}

func (c mockClientSuccess) Info(_ context.Context, _ string, _ ...Field) {}

func (c mockClientSuccess) Notice(_ context.Context, _ string, _ ...Field) {}

func (c mockClientSuccess) Warn(_ context.Context, _ string, _ ...Field) {}

func (c mockClientSuccess) Error(_ context.Context, _ string, _ ...Field) {}

func (c mockClientSuccess) Fatal(_ context.Context, _ string, _ ...Field) {}
