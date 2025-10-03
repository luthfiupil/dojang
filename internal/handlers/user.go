package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luthfiupil/dojang/internal/models"
	"github.com/luthfiupil/dojang/internal/repository"
)

type UserHandler struct {
	Repo *repository.UserRepo
}

func NewUserHandler(db *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		Repo: repository.NewUserRepo(db),
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input models.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	user, err := h.Repo.CreateUser(r.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "invalid date_of_birth") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Default pagination values
	page := 1
	limit := 10

	// Read query params
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	users, err := h.Repo.GetUsers(r.Context(), page, limit)
	if err != nil {
		http.Error(w, "failed to fetch users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
