package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandler(mux *http.ServeMux) {
	handler := &AuthHandler{}
	mux.HandleFunc("POST /auth/register", handler.register())
	mux.HandleFunc("POST /auth/login", handler.login())
}

func (a *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register")
	}
}

func (a *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("login")
	}
}
