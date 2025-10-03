package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luthfiupil/dojang/internal/models"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	query := `
        INSERT INTO users (full_name, email, role_id, date_of_birth, phone, address)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, full_name, email, role_id, date_of_birth, phone, address, created_at
    `

	var user models.User
	var dob *time.Time
	if input.DateOfBirth != nil && !input.DateOfBirth.IsZero() {
		dob = &input.DateOfBirth.Time // ✅ extract the real time.Time
	}

	err := r.DB.QueryRow(ctx, query,
		input.FullName,
		input.Email,
		input.RoleID,
		dob, // ✅ pass *time.Time, not CustomDate
		input.Phone,
		input.Address,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.RoleID,
		&dob, // scan into *time.Time
		&user.Phone,
		&user.Address,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	if dob != nil {
		user.DateOfBirth = &models.CustomDate{Time: *dob}
	}

	return &user, nil
}

func (r *UserRepo) GetUsers(ctx context.Context, page, limit int) ([]models.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := `
        SELECT id, full_name, email, role_id, date_of_birth, phone, address, created_at
        FROM users
        ORDER BY id
        LIMIT $1 OFFSET $2
    `

	rows, err := r.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		var dob *time.Time
		err := rows.Scan(&u.ID, &u.FullName, &u.Email, &u.RoleID, &dob, &u.Phone, &u.Address, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		if dob != nil {
			u.DateOfBirth = &models.CustomDate{Time: *dob}
		}
		users = append(users, u)
	}

	return users, nil
}
