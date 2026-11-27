package response

import (
	"github.com/gin-gonic/gin"
	appErrors "github.com/toby-minigames/server/internal/shared/errors"
)

// EnvelopeSuccess representa la estructura unificada de respuesta exitosa.
type EnvelopeSuccess struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// ErrorBody representa el detalle del error en la respuesta JSON.
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details"`
}

// EnvelopeError representa la estructura unificada de respuesta de error.
type EnvelopeError struct {
	Success bool      `json:"success"`
	Error   ErrorBody `json:"error"`
}

// Success envia una respuesta JSON exitosa con el formato estandar de la plataforma.
func Success(c *gin.Context, status int, data any, message string) {
	c.JSON(status, EnvelopeSuccess{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// Error mapea un error a su codigo HTTP y envia la respuesta JSON de error estandar.
func Error(c *gin.Context, err error) {
	appErr := appErrors.FromError(err)
	c.JSON(appErr.HTTPStatus, EnvelopeError{
		Success: false,
		Error: ErrorBody{
			Code:    appErr.Code,
			Message: appErr.Message,
			Details: appErr.Details,
		},
	})
}
