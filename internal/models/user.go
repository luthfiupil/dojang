package models

import "time"

type User struct {
	ID          int         `json:"id"`
	FullName    string      `json:"full_name"`
	Email       string      `json:"email"`
	RoleID      int         `json:"role_id"`
	DateOfBirth *CustomDate `json:"date_of_birth,omitempty"`
	Phone       *string     `json:"phone,omitempty"`
	Address     *string     `json:"address,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
}

type CreateUserInput struct {
	FullName    string      `json:"full_name"`
	Email       string      `json:"email"`
	RoleID      int         `json:"role_id"`
	DateOfBirth *CustomDate `json:"date_of_birth,omitempty"`
	Phone       *string     `json:"phone,omitempty"`
	Address     *string     `json:"address,omitempty"`
}
