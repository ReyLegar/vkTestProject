package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

func (h *Handler) AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRole != "Admin" {
		http.Error(w, "Only administrators can add movies", http.StatusUnauthorized)
		return
	}

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	movieID, err := h.service.AddMovie(movie)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add movie: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie added successfully. ID: %d\n", movieID)
}

// addMovie обрабатывает запрос на добавление нового фильма.
// @Summary Добавление фильма
// @Description Добавляет новый фильм в систему.
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Информация о фильме"
// @Security ApiKeyAuth
// @Success 201 "Фильм успешно добавлен"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 401 {string} string "Требуется авторизация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movie/add [post]
func (h *Handler) addMovie(w http.ResponseWriter, r *http.Request) {
	h.UserIdentityMiddleware(http.HandlerFunc(h.AddMovieHandler)).ServeHTTP(w, r)
}

func (h *Handler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRole != "Admin" {
		http.Error(w, "Only administrators can update movies", http.StatusUnauthorized)
		return
	}

	queryID := r.URL.Query().Get("id")
	movieID, err := strconv.Atoi(queryID)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateMovie(movieID, movie); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update movie: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Movie updated successfully. ID: %d\n", movieID)
}

// UpdateMovie обрабатывает запрос на обновление информации о фильме.
// @Summary Обновление фильма
// @Description Обновляет информацию о существующем фильме в системе.
// @Tags movies
// @Accept json
// @Produce json
// @Param id query integer true "ID фильма"
// @Param movie body models.Movie true "Информация о фильме"
// @Security ApiKeyAuth
// @Success 200 "Фильм успешно обновлен"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 401 {string} string "Требуется авторизация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movie/update [put]
func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	h.UserIdentityMiddleware(http.HandlerFunc(h.UpdateMovieHandler)).ServeHTTP(w, r)
}

func (h *Handler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRole != "Admin" {
		http.Error(w, "Only administrators can delete movies", http.StatusUnauthorized)
		return
	}

	movieID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteMovie(movieID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete movie: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Movie deleted successfully. ID: %d\n", movieID)
}

// deleteMovie обрабатывает запрос на удаление фильма.
// @Summary Удаление фильма
// @Description Удаляет фильм из системы.
// @Tags movies
// @Param id query integer true "ID фильма"
// @Security ApiKeyAuth
// @Success 200 "Фильм успешно удален"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 401 {string} string "Требуется авторизация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movie/delete [delete]
func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	h.UserIdentityMiddleware(http.HandlerFunc(h.DeleteMovieHandler)).ServeHTTP(w, r)
}

// getAllMovies обрабатывает запрос на получение всех фильмов.
// @Summary Получение всех фильмов
// @Description Получает список всех фильмов в системе.
// @Tags movies
// @Produce json
// @Param sort query string false "Сортировка по алфавиту (asc/desc)"
// @Success 200 {array} models.Movie "Список фильмов"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movie/getAll [get]
func (h *Handler) getAllMovies(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")
	sortBy = strings.ToLower(sortBy)

	movies, err := h.service.GetAllMovies(sortBy)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch movies: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// searchMoviesByTitle обрабатывает запрос на поиск фильмов по названию.
// @Summary Поиск фильмов по названию
// @Description Выполняет поиск фильмов по их названию.
// @Tags movies
// @Param title query string true "Название фильма"
// @Produce json
// @Success 200 {array} models.Movie "Список найденных фильмов"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies/searchByTitle [get]
func (h *Handler) searchMoviesByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := h.service.SearchMoviesByTitle(title)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search movies by title: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// searchMoviesByActorName обрабатывает запрос на поиск фильмов по имени актера.
// @Summary Поиск фильмов по имени актера
// @Description Выполняет поиск фильмов по имени актера, участвующего в них.
// @Tags movies
// @Param actorName query string true "Имя актера"
// @Produce json
// @Success 200 {array} models.Movie "Список найденных фильмов"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies/searchByActorName [get]
func (h *Handler) searchMoviesByActorName(w http.ResponseWriter, r *http.Request) {
	actorName := r.URL.Query().Get("actorName")
	if actorName == "" {
		http.Error(w, "Actor name parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := h.service.SearchMoviesByActorName(actorName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search movies by actor name: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Println(movies)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
