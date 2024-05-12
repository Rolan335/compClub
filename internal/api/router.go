package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface{
	RegisterAdmin(w http.ResponseWriter, r *http.Request)
	NewRent(w http.ResponseWriter, r *http.Request)
	AddNewPC(w http.ResponseWriter, r *http.Request)
	AddNewShift(w http.ResponseWriter, r *http.Request)
	GetAvailablePc(w http.ResponseWriter, r *http.Request)
}

func NewRouter(router Router) http.Handler{
	r := chi.NewRouter()

	r.Post("/registerAdmin", router.RegisterAdmin)
	r.Post("/newRent", router.NewRent)
	r.Post("/addNewPC", router.AddNewPC)
	r.Post("/addNewShift", router.AddNewShift)

	r.Get("/getAvailablePc", router.GetAvailablePc)

	return r
}