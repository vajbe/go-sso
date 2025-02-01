package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-sso/internal/db"
	res "go-sso/internal/middleware"
	"go-sso/internal/types"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "805673868908-mg6cgpd43qcrorj0b2cm7r8iteknkssv.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/oauth/callback",
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint:     google.Endpoint,
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.AddUser(newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been added successfully.", resp)
}

func (h *UserHandler) SSOHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("random-state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *UserHandler) AuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)
	fmt.Fprintf(w, "User Info: %+v", userInfo)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin types.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.Login(userLogin)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to login: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been logged in successfully.", resp)
}
