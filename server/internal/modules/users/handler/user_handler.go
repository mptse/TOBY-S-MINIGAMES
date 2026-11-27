package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/toby-minigames/server/internal/modules/users/application"
	"github.com/toby-minigames/server/internal/modules/users/handler/dto"
	apperrors "github.com/toby-minigames/server/internal/shared/errors"
	"github.com/toby-minigames/server/internal/shared/middleware"
	"github.com/toby-minigames/server/internal/shared/response"
)

// UserHandler expone los endpoints HTTP protegidos del modulo users.
type UserHandler struct {
	service *application.UserService
}

// NewUserHandler crea una nueva instancia de UserHandler.
func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterRoutes registra las rutas protegidas de usuario bajo /users.
func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup, authMW, playerOnlyMW gin.HandlerFunc) {
	usersGroup := rg.Group("/users")
	usersGroup.Use(authMW, playerOnlyMW)
	usersGroup.GET("/me", h.GetMe)
	usersGroup.PATCH("/me", h.UpdateMe)
	usersGroup.GET("/me/games/:slug", h.GetMyGameProfile)
}

func parsePlayerID(c *gin.Context) (uuid.UUID, bool) {
	claims, ok := middleware.GetClaims(c)
	if !ok {
		response.Error(c, apperrors.Unauthorized("missing authentication claims"))
		return uuid.Nil, false
	}

	playerID, err := uuid.Parse(claims.Subject)
	if err != nil {
		response.Error(c, apperrors.Unauthorized("invalid player subject in token"))
		return uuid.Nil, false
	}

	return playerID, true
}

// GetMe maneja GET /api/v1/users/me.
func (h *UserHandler) GetMe(c *gin.Context) {
	playerID, ok := parsePlayerID(c)
	if !ok {
		return
	}

	player, err := h.service.GetProfile(c.Request.Context(), playerID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, dto.FromDomainPlayer(player), "Player profile retrieved successfully")
}

// UpdateMe maneja PATCH /api/v1/users/me.
func (h *UserHandler) UpdateMe(c *gin.Context) {
	playerID, ok := parsePlayerID(c)
	if !ok {
		return
	}

	var req dto.UpdateMeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperrors.Validation("invalid request payload", err.Error()))
		return
	}

	player, err := h.service.UpdateProfile(c.Request.Context(), playerID, req.Nickname, req.Bio, req.AvatarPath)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, dto.FromDomainPlayer(player), "Player profile updated successfully")
}

// GetMyGameProfile maneja GET /api/v1/users/me/games/:slug.
func (h *UserHandler) GetMyGameProfile(c *gin.Context) {
	playerID, ok := parsePlayerID(c)
	if !ok {
		return
	}

	slug := c.Param("slug")
	profile, err := h.service.GetGameProfile(c.Request.Context(), playerID, slug)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, dto.FromDomainGameProfile(profile), "Game profile retrieved successfully")
}
