package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims define los claims personalizados incluidos en los tokens JWT.
type TokenClaims struct {
	Role     string `json:"role"`
	IsGuest  bool   `json:"is_guest"`
	Nickname string `json:"nickname,omitempty"`
	jwt.RegisteredClaims
}

// JWTManager administra la emision y validacion de tokens JWT firmados con HS256.
type JWTManager struct {
	secret     []byte
	expiration time.Duration
}

// NewJWTManager crea una nueva instancia de JWTManager.
func NewJWTManager(secret string, expirationHours int) (*JWTManager, error) {
	if secret == "" {
		return nil, errors.New("jwt secret cannot be empty")
	}
	if expirationHours <= 0 {
		expirationHours = 24
	}
	return &JWTManager{
		secret:     []byte(secret),
		expiration: time.Duration(expirationHours) * time.Hour,
	}, nil
}

// GeneratePlayerToken emite un JWT para un jugador registrado.
func (m *JWTManager) GeneratePlayerToken(playerID, role string) (string, error) {
	return m.generate(playerID, role, false, "")
}

// GenerateGuestToken emite un JWT temporal para un jugador invitado sin persistencia en BD.
func (m *JWTManager) GenerateGuestToken(guestID, nickname string) (string, error) {
	return m.generate(guestID, "guest", true, nickname)
}

func (m *JWTManager) generate(subject, role string, isGuest bool, nickname string) (string, error) {
	now := time.Now().UTC()
	claims := TokenClaims{
		Role:     role,
		IsGuest:  isGuest,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.expiration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}
	return signed, nil
}

// VerifyToken valida la firma y expiracion de un JWT y retorna sus claims.
func (m *JWTManager) VerifyToken(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid or expired token: %w", err)
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
