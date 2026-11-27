package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toby-minigames/server/internal/modules/auth/application"
	"github.com/toby-minigames/server/internal/modules/auth/handler/dto"
	apperrors "github.com/toby-minigames/server/internal/shared/errors"
	"github.com/toby-minigames/server/internal/shared/response"
)

// AuthHandler expone los endpoints HTTP del modulo auth.
type AuthHandler struct {
	service *application.AuthService
}

// NewAuthHandler crea una nueva instancia de AuthHandler.
func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// RegisterRoutes registra las rutas publicas de autenticacion bajo /auth.
func (h *AuthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
	authGroup.POST("/guest", h.Guest)
}

// Register maneja POST /api/v1/auth/register.
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperrors.Validation("invalid request payload", err.Error()))
		return
	}

	token, player, err := h.service.Register(c.Request.Context(), req.Email, req.Password, req.Nickname)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, gin.H{
		"token":  token,
		"player": dto.FromDomainPlayerSummary(player),
	}, "Player registered successfully")
}

// Login maneja POST /api/v1/auth/login.
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperrors.Validation("invalid request payload", err.Error()))
		return
	}

	token, player, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, gin.H{
		"token":  token,
		"player": dto.FromDomainPlayerSummary(player),
	}, "Login successful")
}

// Guest maneja POST /api/v1/auth/guest.
func (h *AuthHandler) Guest(c *gin.Context) {
	var req dto.GuestRequest
	if err := c.ShouldBindJSON(&req); err != nil && !errors.Is(err, io.EOF) {
		response.Error(c, apperrors.Validation("invalid request payload", err.Error()))
		return
	}

	token, guest, err := h.service.Guest(req.Nickname)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, gin.H{
		"token": token,
		"guest": gin.H{
			"guest_id": guest.GuestID,
			"nickname": guest.Nickname,
			"role":     guest.Role,
		},
	}, "Guest session created successfully")
}
