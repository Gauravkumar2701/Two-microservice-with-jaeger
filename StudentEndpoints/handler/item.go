package handler

import (
	"StudentEndpoints/models"
	"StudentEndpoints/tracing"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

var itemIDKey = "itemID"

func student(router chi.Router) {
	router.Get("/", getAllStudent)
	router.Post("/", createStudent)
}

func getAllStudent(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest(gTracer, r)
	defer span.Finish()
	student, err := dbInstance.GetAllStudent()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, student); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}