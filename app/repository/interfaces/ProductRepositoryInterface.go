package interfaces

import (
	"gitlab.com/cinco/app/model"
)

type ProductRepositoryInterface interface {
	ProductInsert(params model.Product) error
}
