package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"github.com/toby-minigames/server/internal/modules/games/domain"
	"gorm.io/gorm"
)

// PostgresCategoryRepository implementa domain.CategoryRepository utilizando GORM y PostgreSQL.
type PostgresCategoryRepository struct {
	db *gorm.DB
}

// NewPostgresCategoryRepository crea una nueva instancia del repositorio de categorias.
func NewPostgresCategoryRepository(db *gorm.DB) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{db: db}
}

// Create inserta una nueva categoria en la base de datos y asigna su ID autogenerado.
func (r *PostgresCategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	model := FromDomainCategory(*category)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	category.ID = model.CategoryID
	return nil
}

// FindBySlug busca una categoria por su slug unico.
func (r *PostgresCategoryRepository) FindBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	var model CategoryModel
	if err := r.db.WithContext(ctx).Where("category_slug = ?", slug).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %v", domain.ErrCategoryNotFound, err)
		}
		return nil, fmt.Errorf("failed to find category by slug: %w", err)
	}
	cat := model.ToDomain()
	return &cat, nil
}

// ListAll retorna todas las categorias ordenadas alfabeticamente por name_category.
func (r *PostgresCategoryRepository) ListAll(ctx context.Context) ([]*domain.Category, error) {
	var models []CategoryModel
	if err := r.db.WithContext(ctx).Order("name_category ASC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}

	categories := make([]*domain.Category, len(models))
	for i := range models {
		cat := models[i].ToDomain()
		categories[i] = &cat
	}
	return categories, nil
}
