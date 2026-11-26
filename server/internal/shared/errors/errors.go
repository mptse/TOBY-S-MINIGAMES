package apperrors

import (
	"errors"
	"net/http"
)

// Codigos estandar de error expuestos en las respuestas JSON.
const (
	CodeValidation   = "VALIDATION_ERROR"
	CodeUnauthorized = "UNAUTHORIZED"
	CodeForbidden    = "FORBIDDEN"
	CodeNotFound     = "NOT_FOUND"
	CodeConflict     = "CONFLICT"
	CodeInternal     = "INTERNAL_ERROR"
)

// AppError representa un error estructurado de aplicacion mapeable a HTTP.
type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
	Details    any
}

// Error implementa la interfaz error de Go.
func (e *AppError) Error() string {
	return e.Message
}

// Validation crea un error HTTP 400 para fallos de validacion de entrada.
func Validation(message string, details any) *AppError {
	return &AppError{
		Code:       CodeValidation,
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
		Details:    details,
	}
}

// Unauthorized crea un error HTTP 401 para credenciales invalidas o token ausente/expirado.
func Unauthorized(message string) *AppError {
	return &AppError{
		Code:       CodeUnauthorized,
		Message:    message,
		HTTPStatus: http.StatusUnauthorized,
	}
}

// Forbidden crea un error HTTP 403 cuando el rol no tiene permisos suficientes.
func Forbidden(message string) *AppError {
	return &AppError{
		Code:       CodeForbidden,
		Message:    message,
		HTTPStatus: http.StatusForbidden,
	}
}

// NotFound crea un error HTTP 404 cuando el recurso solicitado no existe.
func NotFound(message string) *AppError {
	return &AppError{
		Code:       CodeNotFound,
		Message:    message,
		HTTPStatus: http.StatusNotFound,
	}
}

// Conflict crea un error HTTP 409 para conflictos de unicidad (ej. email duplicado).
func Conflict(message string) *AppError {
	return &AppError{
		Code:       CodeConflict,
		Message:    message,
		HTTPStatus: http.StatusConflict,
	}
}

// Internal crea un error HTTP 500 para fallos inesperados del servidor.
func Internal(message string) *AppError {
	return &AppError{
		Code:       CodeInternal,
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
	}
}

// FromError convierte cualquier error en un AppError seguro para respuesta HTTP.
func FromError(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}
	return Internal("unexpected internal server error")
}
