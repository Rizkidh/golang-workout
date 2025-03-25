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

type RepositoryHandler struct {
	usecase *usecase.RepositoryUsecase
}

func NewRepositoryHandler(u *usecase.RepositoryUsecase) *RepositoryHandler {
	return &RepositoryHandler{u}
}

// Create Repository
func (h *RepositoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var repo entity.Repository
	if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := h.usecase.CreateRepository(ctx, &repo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(repo)
}

// Get All Repositories
func (h *RepositoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	repos, err := h.usecase.GetAllRepositories(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch repositories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(repos)
}

// Get Repository by ID
func (h *RepositoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid repository ID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	repo, err := h.usecase.GetRepository(ctx, id)
	if err != nil {
		http.Error(w, "Repository not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(repo)
}

// Update Repository
func (h *RepositoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid repository ID format", http.StatusBadRequest)
		return
	}

	var updatedRepo entity.Repository
	if err := json.NewDecoder(r.Body).Decode(&updatedRepo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedRepo.ID = id
	ctx := context.Background()
	if err := h.usecase.UpdateRepository(ctx, &updatedRepo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Repository updated successfully"})
}

// Delete Repository
func (h *RepositoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid repository ID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := h.usecase.DeleteRepository(ctx, id); err != nil {
		http.Error(w, "Repository not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Repository deleted successfully"})
}
