package dto

import (
	gamesDomain "github.com/toby-minigames/server/internal/modules/games/domain"
)

// GameSummaryResponse representa la vista resumida de un minijuego en el catalogo.
type GameSummaryResponse struct {
	GameID        int    `json:"game_id"`
	Name          string `json:"name"`
	NameSlug      string `json:"name_slug"`
	Description   string `json:"description"`
	ThumbnailSlug string `json:"thumbnail_slug"`
	IsActive      bool   `json:"is_active"`
}

// CategoryResponse representa la vista publica de una categoria.
type CategoryResponse struct {
	CategoryID   int     `json:"category_id"`
	NameCategory string  `json:"name_category"`
	CategoryDesc *string `json:"category_desc"`
	CategorySlug string  `json:"category_slug"`
}

// GameDetailResponse representa el detalle completo de un minijuego incluyendo sus categorias.
type GameDetailResponse struct {
	GameID        int                `json:"game_id"`
	Name          string             `json:"name"`
	NameSlug      string             `json:"name_slug"`
	Description   string             `json:"description"`
	ThumbnailSlug string             `json:"thumbnail_slug"`
	IsActive      bool               `json:"is_active"`
	Categories    []CategoryResponse `json:"categories"`
}

// FromDomainGameSummary convierte una entidad Game de dominio a su resumen DTO.
func FromDomainGameSummary(g *gamesDomain.Game) GameSummaryResponse {
	return GameSummaryResponse{
		GameID:        g.ID,
		Name:          g.Name,
		NameSlug:      g.Slug,
		Description:   g.Description,
		ThumbnailSlug: g.ThumbnailSlug,
		IsActive:      g.IsActive,
	}
}

// FromDomainGameDetail convierte una entidad Game de dominio a su detalle DTO completo.
func FromDomainGameDetail(g *gamesDomain.Game) GameDetailResponse {
	categories := make([]CategoryResponse, len(g.Categories))
	for i, c := range g.Categories {
		categories[i] = CategoryResponse{
			CategoryID:   c.ID,
			NameCategory: c.Name,
			CategoryDesc: c.Description,
			CategorySlug: c.Slug,
		}
	}

	return GameDetailResponse{
		GameID:        g.ID,
		Name:          g.Name,
		NameSlug:      g.Slug,
		Description:   g.Description,
		ThumbnailSlug: g.ThumbnailSlug,
		IsActive:      g.IsActive,
		Categories:    categories,
	}
}
