package slate

import (
	"slate/internal/pkg/http/middleware"

	chi_middleware "github.com/go-chi/chi/middleware"
)

func (api *RESTAPIClient) middleware() {
	api.router.Use(chi_middleware.RequestID)
	api.router.Use(chi_middleware.DefaultCompress)
	api.router.Use(chi_middleware.Logger)
	api.router.Use(chi_middleware.Recoverer)
	api.router.Use(chi_middleware.URLFormat)
	api.router.Use(middleware.PopulateContext)
	api.router.Use(middleware.NewRequestLogger(nil).Info)
}
