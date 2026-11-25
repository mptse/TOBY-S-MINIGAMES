package infrastructure

import (
	"github.com/toby-minigames/server/internal/modules/reviews/domain"
)

// ToDomain convierte un GameReviewModel en una entidad pura de dominio GameReview.
func (m *GameReviewModel) ToDomain() *domain.GameReview {
	return &domain.GameReview{
		ID:        m.ReviewID,
		PlayerID:  m.PlayerID,
		GameID:    m.GameID,
		Rating:    m.Rating,
		Comment:   m.Comment,
		CreatedAt: m.CreatedAt,
	}
}

// FromDomainGameReview convierte una entidad GameReview de dominio en un GameReviewModel de GORM.
func FromDomainGameReview(r *domain.GameReview) *GameReviewModel {
	return &GameReviewModel{
		ReviewID:  r.ID,
		PlayerID:  r.PlayerID,
		GameID:    r.GameID,
		Rating:    r.Rating,
		Comment:   r.Comment,
		CreatedAt: r.CreatedAt,
	}
}
