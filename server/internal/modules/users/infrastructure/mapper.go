package infrastructure

import (
	"github.com/toby-minigames/server/internal/modules/users/domain"
	"github.com/toby-minigames/server/internal/shared/database"
)

// ToDomain convierte un PlayerModel en una entidad pura de dominio Player.
func (m *PlayerModel) ToDomain() *domain.Player {
	return &domain.Player{
		ID:         m.PlayerID,
		Email:      m.Email,
		Password:   m.Password,
		Nickname:   m.Nickname,
		Bio:        m.Bio,
		Role:       domain.PlayerRole(m.Role),
		Exp:        m.Exp,
		Tobycoin:   m.Tobycoin,
		AvatarPath: m.AvatarPath,
	}
}

// FromDomainPlayer convierte una entidad Player de dominio en un PlayerModel de GORM.
func FromDomainPlayer(p *domain.Player) *PlayerModel {
	return &PlayerModel{
		PlayerID:   p.ID,
		Email:      p.Email,
		Password:   p.Password,
		Nickname:   p.Nickname,
		Bio:        p.Bio,
		Role:       string(p.Role),
		Exp:        p.Exp,
		Tobycoin:   p.Tobycoin,
		AvatarPath: p.AvatarPath,
	}
}

// ToDomain convierte un GameProfileModel en una entidad pura de dominio GameProfile.
func (m *GameProfileModel) ToDomain() (*domain.GameProfile, error) {
	dataMap, err := database.JSONBToMap(m.GameData)
	if err != nil {
		return nil, err
	}

	return &domain.GameProfile{
		ID:            m.ProfileID,
		PlayerID:      m.PlayerID,
		GameID:        m.GameID,
		HighestScore:  m.HighestScore,
		MatchesPlayed: m.MatchesPlayed,
		GameData:      dataMap,
		UpdatedAt:     m.UpdatedAt,
	}, nil
}

// FromDomainGameProfile convierte una entidad GameProfile de dominio en un GameProfileModel de GORM.
func FromDomainGameProfile(gp *domain.GameProfile) (*GameProfileModel, error) {
	jsonData, err := database.MapToJSONB(gp.GameData)
	if err != nil {
		return nil, err
	}

	return &GameProfileModel{
		ProfileID:     gp.ID,
		PlayerID:      gp.PlayerID,
		GameID:        gp.GameID,
		HighestScore:  gp.HighestScore,
		MatchesPlayed: gp.MatchesPlayed,
		GameData:      jsonData,
		UpdatedAt:     gp.UpdatedAt,
	}, nil
}
