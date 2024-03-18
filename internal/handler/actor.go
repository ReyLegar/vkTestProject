package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

func (h *Handler) AddActorHandler(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRole != "Admin" {
		http.Error(w, "Only administrators can add actors", http.StatusUnauthorized)
		return
	}

	var actor models.Actor
	err = json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	actorID, err := h.service.ActorRepository.AddActor(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{"actor_id": actorID}
	jsonResponse(w, http.StatusCreated, response)
}

// AddActor добавляет актера в систему.
// @Summary Добавление актера
// @Description Добавляет нового актера в систему.
// @Tags actors
// @Accept json
// @Produce json
// @Param actor body models.Actor true "Информация об актере"
// @Security ApiKeyAuth
// @Success 201 {object} map[string]int "ID нового актера"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 401 {string} string "Требуется авторизация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actor/add [post]
func (h *Handler) AddActor(w http.ResponseWriter, r *http.Request) {
	h.UserIdentityMiddleware(http.HandlerFunc(h.AddActorHandler)).ServeHTTP(w, r)
}

func (h *Handler) UpdateActorHandler(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRole != "Admin" {
		http.Error(w, "Only administrators can update actors", http.StatusUnauthorized)
		return
	}

	actorID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var actor models.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateActor(actorID, actor); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateActor удаляет актера из системы.
// @Summary Обновление актера
// @Description Обновляет актера из системы по его идентификатору.
// @Tags actors
// @Param id query integer true "ID актера"
// @Security ApiKeyAuth
// @Success 204 "Актер успешно удален"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 401 {string} string "Требуется авторизация"
// @Failure 404 {string} string "Актер не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actor/update [delete]
func (h *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	h.UserIdentityMiddleware(http.HandlerFunc(h.UpdateActorHandler)).ServeHTTP(w, r)
}

func (h *Handler) DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRole != "Admin" {
		http.Error(w, "Only administrators can delete actors", http.StatusUnauthorized)
		return
	}

	actorID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteActor(actorID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteActor удаляет актера из системы.
// @Summary Удаление актера
// @Description Удаляет актера из системы по его идентификатору.
// @Tags actors
// @Param id query integer true "ID актера"
// @Security ApiKeyAuth
// @Success 204 "Актер успешно удален"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 401 {string} string "Требуется авторизация"
// @Failure 404 {string} string "Актер не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actor/delete [delete]
func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	h.UserIdentityMiddleware(http.HandlerFunc(h.DeleteActorHandler)).ServeHTTP(w, r)
}

// GetAllActorsAndMovies возвращает всех актеров и фильмы, в которых они снимались.
// @Summary Получение всех актеров и фильмов
// @Description Получает список всех актеров и фильмов, в которых они снимались.
// @Tags actors
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string][]models.Movie "Список актеров и их фильмов"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actor/getAll [get]
func (h *Handler) GetAllActorsAndMovies(w http.ResponseWriter, r *http.Request) {
	actorsAndMovies, err := h.service.GetAllActorsAndMovies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch actors and movies: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actorsAndMovies)
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
