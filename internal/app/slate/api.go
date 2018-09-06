package slate

import (
	"slate/internal/app/slate/route"

	"github.com/go-chi/chi"
)

func (api *RESTAPIClient) api() {
	api.router.Route("/", func(router chi.Router) {
		router.Get("/", route.HelloWorld(api.deps.SecretClient))
	})
}
