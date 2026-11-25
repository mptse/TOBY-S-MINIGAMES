package infrastructure

import (
	"github.com/toby-minigames/server/internal/modules/scores/domain"
)

// ToDomain convierte un MatchHistoryModel en una entidad pura de dominio MatchHistory.
func (m *MatchHistoryModel) ToDomain() *domain.MatchHistory {
	return &domain.MatchHistory{
		ID:              m.MatchID,
		PlayerID:        m.PlayerID,
		GameID:          m.GameID,
		GameMode:        m.GameMode,
		ScoreObtained:   m.ScoreObtained,
		XPEarned:        m.XPEarned,
		TobycoinsEarned: m.TobycoinsEarned,
		PlayedAt:        m.PlayedAt,
	}
}

// FromDomainMatchHistory convierte una entidad MatchHistory de dominio en un MatchHistoryModel de GORM.
func FromDomainMatchHistory(mh *domain.MatchHistory) *MatchHistoryModel {
	return &MatchHistoryModel{
		MatchID:         mh.ID,
		PlayerID:        mh.PlayerID,
		GameID:          mh.GameID,
		GameMode:        mh.GameMode,
		ScoreObtained:   mh.ScoreObtained,
		XPEarned:        mh.XPEarned,
		TobycoinsEarned: mh.TobycoinsEarned,
		PlayedAt:        mh.PlayedAt,
	}
}
