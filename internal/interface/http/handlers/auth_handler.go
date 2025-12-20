package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Saad7890-web/internal/application/auth"

	"github.com/Saad7890-web/pkg/response"
)

type AuthHandler struct {
	authService *auth.Service
	jwtService  *auth.JWTService
}

func NewAuthHandler(
	authService *auth.Service,
	jwtService *auth.JWTService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request){
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	u, err :=h.authService.Signup(
		r.Context(),
		req.Email,
		req.Password,
		req.Name,
	)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.jwtService.GenerateToken(u.ID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "token generation failed")
		return
	}

	response.JSON(w, http.StatusCreated, map[string]string {
		"token": token,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	u, err := h.authService.Login(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := h.jwtService.GenerateToken(u.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "token generation failed")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
