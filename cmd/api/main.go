package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hananadira/ecommarce-api/internal/config"
	"github.com/hananadira/ecommarce-api/internal/database"
	"github.com/hananadira/ecommarce-api/internal/handler"
	"github.com/hananadira/ecommarce-api/internal/repository"
	"github.com/hananadira/ecommarce-api/internal/service"
)

func main() {
	cfg := config.Load()

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET belun diisi di .env")
	}

	pool, err := database.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("gagal konek database: %v", err)
	}
	defer pool.Close()
	log.Println("database terhubung")

	userRepo := repository.NewUserRepository(pool)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status":"ok"})
	})

	auth := r.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server gagal jalan: %v", err)
	}
}