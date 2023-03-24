package service

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/repository/interfaces"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
)

type ProductService struct {
	productRepository interfaces.ProductRepositoryInterface
}

func (service ProductService) CreateProduct(ctx *fiber.Ctx, params *param.Request) (int, error) {
	var createdProduct model.Product

	for _, v := range params.Data {
		time.Sleep(10 * time.Millisecond)
		// Insert the record into the database
		createdProduct = model.Product{
			RequestID: params.RequestID,
			ID:        v.ID,
			Customers: v.Customers,
			Quantity:  v.Quantity,
			Price:     v.Price,
			CreatedAt: time.Now(),
		}
	}

	err := service.productRepository.ProductInsert(createdProduct)
	if err != nil {
		log.Printf("Failed to insert record: %v\n", err)
	}

	return fiber.StatusOK, nil
}

func NewProductService(repository interfaces.ProductRepositoryInterface) serviceInterface.ProductServiceInterface {
	return &ProductService{
		productRepository: repository,
	}
}
