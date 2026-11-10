package handler 

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hananadira/ecommarce-api/internal/dto"
	"github.com/hananadira/ecommarce-api/internal/repository"
	"github.com/hananadira/ecommarce-api/internal/service"
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

func (h * AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data tidak valid: " + err.Error()})
	return 
	}

	user, err := h.auth.Register(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, repository.ErrEmailTaken) {
			c.JSON(http.StatusConflict, gin.H{"error": "email sudah terdaftar"})
			return
		}
		log.Printf("register gagal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan pada server"})
		return
	}

	c.JSON(http.StatusCreated, dto.UserResponse{
		ID: 	  user.ID,
		Name: 	  user.Name,
		Email: 	  user.Email,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data tidak valid: " + err.Error()})
		return
	}

	token, user, err := h.auth.Login(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email atau password salah"})
			return
		}
		log.Printf("login gagal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan pada server"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token, 
		User: dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
		},
	})
}