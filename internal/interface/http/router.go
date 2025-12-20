package http

import (
	"net/http"

	"github.com/Saad7890-web/internal/interface/http/handlers"
)

func NewRouter(authHandler *handlers.AuthHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", authHandler.Signup)
	mux.HandleFunc("/login", authHandler.Login)

	return mux
}