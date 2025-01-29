package db

import (
	"context"
	"fmt"
	"go-sso/internal/types"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(user types.User) (types.UserResponse, error) {
	pool := GetDBPool()
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	var id string

	password_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return types.UserResponse{Email: user.Email}, err
	}

	err = pool.QueryRow(context.Background(), query, user.Email, string(password_hash)).Scan(&id)
	if err != nil {
		return types.UserResponse{Email: user.Email}, fmt.Errorf("failed to insert record: %w", err)
	}
	response := types.UserResponse{
		Email: user.Email,
		Id:    id,
	}
	return response, nil
}
