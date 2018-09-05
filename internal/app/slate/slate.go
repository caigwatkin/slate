package slate

import (
	"context"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secret"
)

type Client struct {
	secretClient *secret.Client
}

func NewClient(secretClient *secret.Client) Client {
	return Client{
		secretClient: secretClient,
	}
}

func (c Client) ListenAndServe(ctx context.Context, port int) {
	log.Info(ctx, "hello world", log.FmtInt(port, "port"))
}
