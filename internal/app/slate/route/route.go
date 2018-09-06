package route

import (
	"net/http"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secret"
)

// TODO use secret client
func HelloWorld(_ *secret.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Context(), "hello world")
	}
}
