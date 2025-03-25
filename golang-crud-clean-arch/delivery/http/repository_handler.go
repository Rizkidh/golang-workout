package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang-crud-clean-arch/internal/entity"
	"golang-crud-clean-arch/internal/usecase"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryHandler struct {
	repoUsecase *usecase.RepositoryUsecase
	userUsecase *usecase.UserUsecase // Tambahkan UserUsecase untuk GetUser
}

func NewRepositoryHandler(repoU *usecase.RepositoryUsecase, userU *usecase.UserUsecase) *RepositoryHandler {
	return &RepositoryHandler{
		repoUsecase: repoU,
		userUsecase: userU,
	}
}

// Create Repository
func (h *RepositoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var repo entity.Repository
	if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo.ID = primitive.NewObjectID()
	ctx := context.Background()
	if err := h.repoUsecase.CreateRepository(ctx, &repo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repo)
}

// Get All Repositories
func (h *RepositoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	repos, err := h.repoUsecase.GetAllRepositories(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch repositories", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repos)
}

// Get Repository by ID
func (h *RepositoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid repository ID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	repo, err := h.repoUsecase.GetRepository(ctx, id)
	if err != nil {
		http.Error(w, "Repository not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repo)
}

// GetUser retrieves a user by ObjectID
func (h *RepositoryHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari parameter URL
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID format"}`, http.StatusBadRequest)
		return
	}

	// Gunakan context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ambil data user dari usecase
	_, err = h.userUsecase.GetUser(ctx, id)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	}

	// Set header response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// Update Repository
func (h *RepositoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
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
	if err := h.repoUsecase.UpdateRepository(ctx, &updatedRepo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Repository updated successfully"})
}

// Delete Repository
func (h *RepositoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid repository ID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := h.repoUsecase.DeleteRepository(ctx, id); err != nil {
		http.Error(w, "Repository not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Repository deleted successfully"})
}
