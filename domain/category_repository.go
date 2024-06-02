package domain

type CategoryRepository interface {
	CreateCategory(category Category) (Category, error)
	FindCategoryById(id int) (Category, error)
	FindAllCategories() ([]Category, error)
}
