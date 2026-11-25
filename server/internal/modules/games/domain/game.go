package domain

import (
	"errors"
	"strings"
	"time"
)

// Category representa una etiqueta de clasificacion para los minijuegos.
type Category struct {
	ID          int
	Name        string
	Description *string
	Slug        string
}

// Game representa un minijuego registrado en el catalogo de la plataforma.
type Game struct {
	ID            int
	Name          string
	Slug          string
	Description   string
	ThumbnailSlug string
	IsActive      bool
	CreatedAt     time.Time
	Categories    []Category
}

// NewGame crea y valida una nueva entidad Game.
func NewGame(name, slug, description, thumbnailSlug string) (*Game, error) {
	name = strings.TrimSpace(name)
	slug = strings.TrimSpace(slug)
	description = strings.TrimSpace(description)
	thumbnailSlug = strings.TrimSpace(thumbnailSlug)

	if name == "" {
		return nil, errors.New("game name cannot be empty")
	}
	if slug == "" {
		return nil, errors.New("game slug cannot be empty")
	}
	if description == "" {
		return nil, errors.New("game description cannot be empty")
	}
	if thumbnailSlug == "" {
		return nil, errors.New("game thumbnail slug cannot be empty")
	}

	return &Game{
		Name:          name,
		Slug:          slug,
		Description:   description,
		ThumbnailSlug: thumbnailSlug,
		IsActive:      true,
		CreatedAt:     time.Now().UTC(),
		Categories:    []Category{},
	}, nil
}

// NewCategory crea y valida una nueva entidad Category.
func NewCategory(name, slug string, description *string) (*Category, error) {
	name = strings.TrimSpace(name)
	slug = strings.TrimSpace(slug)

	if name == "" {
		return nil, errors.New("category name cannot be empty")
	}
	if slug == "" {
		return nil, errors.New("category slug cannot be empty")
	}

	return &Category{
		Name:        name,
		Description: description,
		Slug:        slug,
	}, nil
}
