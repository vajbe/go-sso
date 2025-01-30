package db

import (
	"context"
	"fmt"
	"go-sso/internal/types"
	"go-sso/internal/utils"
	"time"

	"github.com/jackc/pgx/v5"
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

func Login(user types.UserLogin) (types.UserLoginResponse, error) {

	isValid, err := verifyUser(user)

	if err != nil {
		return types.UserLoginResponse{Email: user.Email}, err
	}
	if !isValid {
		return types.UserLoginResponse{Email: user.Email}, fmt.Errorf("invalid credentials")
	}

	// Generate JWT
	jwtToken, err := utils.GenerateJwtToken(user.Email)
	if err != nil {
		return types.UserLoginResponse{Email: user.Email}, err
	}

	fmt.Printf("Token generated is %s\n", jwtToken)

	// Generate Redis key
	key := fmt.Sprintf("jwt:%s", jwtToken)
	err = REDIS_CLIENT.Set(context.Background(), key, jwtToken, 10*time.Second).Err()
	if err != nil {
		return types.UserLoginResponse{Email: user.Email}, err
	}
	return types.UserLoginResponse{Email: user.Email}, nil
}

func verifyUser(user types.UserLogin) (bool, error) {
	pool := GetDBPool()
	var password_hash string

	err := pool.QueryRow(context.Background(), `SELECT PASSWORD_HASH FROM USERS WHERE EMAIL=$1`, user.Email).Scan(&password_hash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("invalid credentials")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(user.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("invalid credentials")
		}
		return false, err
	}

	return true, nil
}
