package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	apperrors "github.com/toby-minigames/server/internal/shared/errors"
	"github.com/toby-minigames/server/internal/shared/response"
	"github.com/toby-minigames/server/internal/shared/security"
)

const claimsContextKey = "auth_claims"

// AuthRequired valida el token Bearer JWT y guarda los claims en el contexto de Gin.
func AuthRequired(jwtMgr *security.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, apperrors.Unauthorized("missing or malformed authorization header"))
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtMgr.VerifyToken(strings.TrimSpace(tokenStr))
		if err != nil {
			response.Error(c, apperrors.Unauthorized("invalid or expired token"))
			c.Abort()
			return
		}

		c.Set(claimsContextKey, claims)
		c.Next()
	}
}

// PlayerOnly bloquea el acceso a sesiones de invitados (is_guest = true) con HTTP 403.
func PlayerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := GetClaims(c)
		if !ok {
			response.Error(c, apperrors.Unauthorized("missing authentication claims"))
			c.Abort()
			return
		}

		if claims.IsGuest || claims.Role == "guest" {
			response.Error(c, apperrors.Forbidden("guest accounts cannot access this resource"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetClaims extrae los claims JWT validados desde el contexto de Gin.
func GetClaims(c *gin.Context) (*security.TokenClaims, bool) {
	val, exists := c.Get(claimsContextKey)
	if !exists {
		return nil, false
	}
	claims, ok := val.(*security.TokenClaims)
	return claims, ok
}

// CORS configura las cabeceras de acceso entre origenes para clientes web y mock-client.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
