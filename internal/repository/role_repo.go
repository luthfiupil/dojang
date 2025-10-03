package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luthfiupil/dojang/internal/models"
)

type RoleRepo struct {
	DB *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) *RoleRepo {
	return &RoleRepo{DB: db}
}

func (r *RoleRepo) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	rows, err := r.DB.Query(ctx, `SELECT id, role_name FROM roles ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.RoleName); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
