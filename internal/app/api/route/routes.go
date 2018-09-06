package route

import (
	"net/http"
	"slate/internal/pkg/log"
)

func HelloWorld() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Context(), "hello world")
	}
}
