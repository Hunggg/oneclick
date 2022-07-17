package categories

import (
	"oneclick/entity"
) 

type PersistentDB interface {
	GetCategoryById(id uint64) (entity.Categories, error)
	GetListCategory(offset int, limit int) ([]entity.Categories, uint64, error)

	UpdateCategoryById(id uint64) (entity.Categories, error)
	SaveCategory(category entity.Categories) error
	SaveBatchCategory(listCategory []entity.Categories) error

	DeleteCategoryById(id uint64) error
}