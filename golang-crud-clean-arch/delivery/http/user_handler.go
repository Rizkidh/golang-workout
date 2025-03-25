package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"golang-crud-clean-arch/m/internal/entity"
	"golang-crud-clean-arch/m/internal/usecase"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := h.usecase.CreateUser(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	user, err := h.usecase.GetUser(ctx, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	var updatedUser entity.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser.ID = id
	ctx := context.Background()
	if err := h.usecase.UpdateUser(ctx, &updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := h.usecase.DeleteUser(ctx, id); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
