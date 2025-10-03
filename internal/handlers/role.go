package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luthfiupil/dojang/internal/repository"
)

type RoleHandler struct {
	Repo *repository.RoleRepo
}

func NewRoleHandler(db *pgxpool.Pool) *RoleHandler {
	return &RoleHandler{
		Repo: repository.NewRoleRepo(db),
	}
}

func (h *RoleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.Repo.GetAllRoles(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch roles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}
