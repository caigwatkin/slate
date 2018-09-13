package slate

import (
	"slate/internal/app/slate/routes"

	"github.com/go-chi/chi"
)

func (api *RESTAPIClient) api() {
	api.router.Route("/", func(router chi.Router) {
		router.Get("/", routes.Ping())
	})
}
