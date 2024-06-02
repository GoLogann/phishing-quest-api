package repository

import "phishing-quest/internal/domain/entity"

type CategoryRepository interface {
	CreateCategory(category entity.Category) (entity.Category, error)
	FindCategoryById(id int) (entity.Category, error)
	FindAllCategories() ([]entity.Category, error)
}
