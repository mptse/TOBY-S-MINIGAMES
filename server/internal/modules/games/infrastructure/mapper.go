package infrastructure

import (
	"github.com/toby-minigames/server/internal/modules/games/domain"
)

// ToDomain convierte un CategoryModel en una entidad pura de dominio Category.
func (m *CategoryModel) ToDomain() domain.Category {
	return domain.Category{
		ID:          m.CategoryID,
		Name:        m.NameCategory,
		Description: m.CategoryDesc,
		Slug:        m.CategorySlug,
	}
}

// FromDomainCategory convierte una entidad Category de dominio en un CategoryModel de GORM.
func FromDomainCategory(c domain.Category) CategoryModel {
	return CategoryModel{
		CategoryID:   c.ID,
		NameCategory: c.Name,
		CategoryDesc: c.Description,
		CategorySlug: c.Slug,
	}
}

// ToDomain convierte un GameModel en una entidad pura de dominio Game.
func (m *GameModel) ToDomain() *domain.Game {
	categories := make([]domain.Category, len(m.Categories))
	for i, c := range m.Categories {
		categories[i] = c.ToDomain()
	}

	return &domain.Game{
		ID:            m.GameID,
		Name:          m.Name,
		Slug:          m.NameSlug,
		Description:   m.Description,
		ThumbnailSlug: m.ThumbnailSlug,
		IsActive:      m.IsActive,
		CreatedAt:     m.CreateAt,
		Categories:    categories,
	}
}

// FromDomainGame convierte una entidad Game de dominio en un GameModel de GORM.
func FromDomainGame(g *domain.Game) *GameModel {
	categories := make([]CategoryModel, len(g.Categories))
	for i, c := range g.Categories {
		categories[i] = FromDomainCategory(c)
	}

	return &GameModel{
		GameID:        g.ID,
		Name:          g.Name,
		NameSlug:      g.Slug,
		Description:   g.Description,
		ThumbnailSlug: g.ThumbnailSlug,
		IsActive:      g.IsActive,
		CreateAt:      g.CreatedAt,
		Categories:    categories,
	}
}
