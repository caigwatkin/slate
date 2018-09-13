package routes

import (
	"encoding/json"
	"net/http"
	"slate/internal/pkg/http/render"
	"slate/internal/pkg/log"
)

func Ping() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Info(ctx, "Pinging")

		b, err := json.MarshalIndent(map[string]interface{}{
			"hello": "world",
		}, "", "\t")
		if err != nil {
			render.ErrorOrStatus(ctx, w, err)
		}
		render.ContentJSON(ctx, w, b)
	}
}
