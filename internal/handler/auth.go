package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

// SignUpHandler обрабатывает запрос на регистрацию нового пользователя.
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя в системе.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Информация о пользователе"
// @Success 200 {object} map[string]int "ID нового пользователя"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /auth/signUp [post]
func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// SignInHandler обрабатывает запрос на аутентификацию пользователя.
// @Summary Аутентификация пользователя
// @Description Аутентифицирует пользователя в системе.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Информация о пользователе"
// @Success 200 {object} map[string]int "Токен аутентификации"
// @Failure 400 {string} string "Некорректные данные запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /auth/signIn [post]
func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var input models.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.PasswordHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"token": token})
}
