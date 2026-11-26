package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword genera un hash bcrypt seguro a partir de una contrasena en texto plano.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// CheckPassword compara una contrasena en texto plano contra su hash bcrypt almacenado.
func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
