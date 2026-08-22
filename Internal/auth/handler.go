package auth

import (
	"brief-url/configs"
	"brief-url/pkg/response"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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
		fmt.Println("register")
	}
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Email == "" || req.Password == "" {
			response.Json(w, "email or password is empty", http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := LoginResponse{
			Token: "abc",
		}

		response.Json(w, res, http.StatusOK)
	}
}
