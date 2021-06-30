package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/opentracing/opentracing-go"
	"net/http"
)
var gTracer opentracing.Tracer
func NewHandler(tracer opentracing.Tracer) http.Handler {
	router := chi.NewRouter()
	gTracer=tracer
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	router.Route("/student", student)

	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)

}