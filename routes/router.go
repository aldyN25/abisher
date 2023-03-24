package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/cinco/app/repository"
	"gorm.io/gorm"

	"gitlab.com/cinco/app/handler"
	"gitlab.com/cinco/app/service"
)

func AllRouter(app *fiber.App, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)

	productService := service.NewProductService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	//Handler := handler.NewHandler(services)
	api := app.Group("/api", logger.New())

	api.Post("/product", productHandler.CreateProduct)
}
