package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toby-minigames/server/internal/modules/games/application"
	"github.com/toby-minigames/server/internal/modules/games/handler/dto"
	"github.com/toby-minigames/server/internal/shared/response"
)

// GameHandler expone los endpoints HTTP publicos del catalogo de minijuegos.
type GameHandler struct {
	service *application.GameService
}

// NewGameHandler crea una nueva instancia de GameHandler.
func NewGameHandler(service *application.GameService) *GameHandler {
	return &GameHandler{service: service}
}

// RegisterRoutes registra las rutas publicas de juegos bajo /games.
func (h *GameHandler) RegisterRoutes(rg *gin.RouterGroup) {
	gamesGroup := rg.Group("/games")
	gamesGroup.GET("", h.List)
	gamesGroup.GET("/:slug", h.GetBySlug)
}

// List maneja GET /api/v1/games.
func (h *GameHandler) List(c *gin.Context) {
	categorySlug := c.Query("category")
	games, err := h.service.ListGames(c.Request.Context(), categorySlug)
	if err != nil {
		response.Error(c, err)
		return
	}

	items := make([]dto.GameSummaryResponse, len(games))
	for i, g := range games {
		items[i] = dto.FromDomainGameSummary(g)
	}

	response.Success(c, http.StatusOK, gin.H{"games": items}, "Games retrieved successfully")
}

// GetBySlug maneja GET /api/v1/games/:slug.
func (h *GameHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	game, err := h.service.GetGameBySlug(c.Request.Context(), slug)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, dto.FromDomainGameDetail(game), "Game detail retrieved successfully")
}
