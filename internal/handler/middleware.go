package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "userRole"
)

func (h *Handler) UserIdentityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		if len(headerParts[1]) == 0 {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}

		userId, err := h.service.Authorization.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		role, err := h.service.Authorization.GetRoleFromToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)
		ctx = context.WithValue(ctx, roleCtx, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserId(r *http.Request) (int, error) {
	id := r.Context().Value(userCtx)
	if id == nil {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func GetUserRole(r *http.Request) (string, error) {
	role := r.Context().Value(roleCtx)
	if role == nil {
		return "", errors.New("user role not found")
	}

	roleStr, ok := role.(string)
	if !ok {
		return "", errors.New("user role is of invalid type")
	}

	return roleStr, nil
}
