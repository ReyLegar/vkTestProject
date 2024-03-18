package handler

import (
	"net/http"

	"github.com/ReyLegar/vkTestProject/internal/service"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/movie/add":
		h.addMovie(w, r)
	case "/movie/update":
		h.updateMovie(w, r)
	case "/movie/delete":
		h.deleteMovie(w, r)
	case "/movie/getAll":
		h.getAllMovies(w, r)
	case "/movie/searchByTitle":
		h.searchMoviesByTitle(w, r)
	case "/movie/searchByActorName":
		h.searchMoviesByActorName(w, r)
	case "/actor/add":
		h.AddActor(w, r)
	case "/actor/update":
		h.UpdateActor(w, r)
	case "/actor/delete":
		h.DeleteActor(w, r)
	case "/actor/getAll":
		h.GetAllActorsAndMovies(w, r)
	case "/auth/signUp":
		h.SignUpHandler(w, r)
	case "/auth/signIn":
		h.SignInHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

/* func (h *Handler) InitRoutes() {

	http.HandleFunc("/auth/signUp", h.SignUpHandler)
	http.HandleFunc("/auth/signIn", h.SignInHandler)

	http.HandleFunc("/movie/add", h.addMovie)
	http.HandleFunc("/movie/update/", h.updateMovie)
	http.HandleFunc("/movie/delete/", h.deleteMovie)
	http.HandleFunc("/movie/getAll", h.getAllMovies)
	http.HandleFunc("/movie/searchByTitle", h.searchMoviesByTitle)
	http.HandleFunc("/movie/searchByActorName", h.searchMoviesByActorName)

	http.HandleFunc("/actor/add", h.AddActor)
	http.HandleFunc("/actor/update", h.UpdateActor)
	http.HandleFunc("/actor/delete", h.DeleteActor)
} */
