package repository

import (
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/repository/interfaces"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func (r ProductRepository) ProductInsert(params model.Product) error {
	err := r.Db.Debug().Table("public.products").Create(&params).Error
	return err
}
func NewProductRepository(db *gorm.DB) interfaces.ProductRepositoryInterface {
	return &ProductRepository{
		Db: db,
	}
}
