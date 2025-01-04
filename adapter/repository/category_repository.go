package repository

import (
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type ICategoryRepository interface {
	IRepository[domain.Category]
	GetByCategoryName(categoryName string) (*domain.Category, error)
}

type CategoryRepository struct {
	IRepository[domain.Category]
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		IRepository: NewRepository[domain.Category](db),
		db:          db,
	}
}

func (cr *CategoryRepository) GetByCategoryName(categoryName string) (*domain.Category, error) {
	var category domain.Category
	if err := cr.db.Where("category_name = ?", categoryName).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
