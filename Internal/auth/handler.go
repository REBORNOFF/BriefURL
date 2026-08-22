package auth

import (
	"brief-url/configs"
	"brief-url/pkg/request"
	"brief-url/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	Config *configs.Config
}

type AuthHandlerDeps struct {
	Config *configs.Config
}

func NewAuthHandler(mux *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	mux.HandleFunc("POST /auth/register", handler.register())
	mux.HandleFunc("POST /auth/login", handler.login())
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}

		fmt.Println(body)

		res := LoginResponse{
			Token: "abc",
		}

		response.Json(w, res, http.StatusOK)
	}
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		fmt.Println(body)

		res := LoginResponse{
			Token: "abc",
		}

		response.Json(w, res, http.StatusOK)
	}
}
