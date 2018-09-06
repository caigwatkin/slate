package middleware

import (
	"slate/internal/pkg/middleware"

	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
)

func Default(router *chi.Mux) {
	router.Use(chi_middleware.RequestID)
	router.Use(chi_middleware.DefaultCompress)
	router.Use(chi_middleware.Logger)
	router.Use(chi_middleware.Recoverer)
	router.Use(chi_middleware.URLFormat)
	router.Use(middleware.PopulateContext)
	router.Use(middleware.NewRequestLogger(nil).Info)
}
