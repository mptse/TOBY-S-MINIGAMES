package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// PlayerRole representa el rol de autorizacion asignado a un jugador.
type PlayerRole string

const (
	RolePlayer PlayerRole = "player"
	RoleAdmin  PlayerRole = "admin"
)

// Player representa la cuenta principal del usuario y su estado de progresion en la plataforma.
type Player struct {
	ID         uuid.UUID
	Email      string
	Password   string
	Nickname   string
	Bio        *string
	Role       PlayerRole
	Exp        int
	Tobycoin   int
	AvatarPath string
}

// NewPlayer crea y valida una nueva entidad Player con valores iniciales por defecto.
func NewPlayer(email, passwordHash, nickname string) (*Player, error) {
	email = strings.TrimSpace(email)
	nickname = strings.TrimSpace(nickname)

	if email == "" {
		return nil, errors.New("player email cannot be empty")
	}
	if passwordHash == "" {
		return nil, errors.New("player password hash cannot be empty")
	}
	if nickname == "" {
		return nil, errors.New("player nickname cannot be empty")
	}

	return &Player{
		ID:         uuid.New(),
		Email:      email,
		Password:   passwordHash,
		Nickname:   nickname,
		Role:       RolePlayer,
		Exp:        0,
		Tobycoin:   0,
		AvatarPath: "default_avatar.png",
	}, nil
}

// AddRewards incrementa la experiencia y los Tobycoins asegurando valores no negativos.
func (p *Player) AddRewards(expDelta, coinsDelta int) error {
	if expDelta < 0 || coinsDelta < 0 {
		return errors.New("reward deltas cannot be negative")
	}
	p.Exp += expDelta
	p.Tobycoin += coinsDelta
	return nil
}

// DeductCoins descuenta Tobycoins del saldo del jugador si cuenta con fondos suficientes.
func (p *Player) DeductCoins(amount int) error {
	if amount <= 0 {
		return errors.New("deduction amount must be greater than zero")
	}
	if p.Tobycoin < amount {
		return errors.New("insufficient tobycoin balance")
	}
	p.Tobycoin -= amount
	return nil
}
