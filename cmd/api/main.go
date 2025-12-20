package main

import (
	"log"
	"net/http"

	"github.com/Saad7890-web/internal/application/auth"
	jwtAuth "github.com/Saad7890-web/internal/application/auth"
	"github.com/Saad7890-web/internal/config"
	"github.com/Saad7890-web/internal/infrastructure/db"
	"github.com/Saad7890-web/internal/infrastructure/repository"
	httpInterface "github.com/Saad7890-web/internal/interface/http"
	"github.com/Saad7890-web/internal/interface/http/handlers"
)	

func main(){
	cfg := config.LoadConfig()

	dbConn := db.NewPostgresDb(cfg)

	userRepo := repository.NewUserRepository(dbConn)
	authService := auth.NewService(userRepo)
	jwtService := jwtAuth.NewJWTService(cfg.JWTSecret)

	authHandler := handlers.NewAuthHandler(authService, jwtService)

	router := httpInterface.NewRouter(authHandler)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
