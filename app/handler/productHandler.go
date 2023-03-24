package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/service/interfaces"
	utilities "gitlab.com/cinco/utils"
)

type Product interface {
	CreateProduct(ctx *fiber.Ctx) error
}

type ProductHandler struct {
	ProductService interfaces.ProductServiceInterface
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var params *param.Request
	err := ctx.BodyParser(&params)
	if err != nil {
		return ctx.Status(400).
			JSON(fiber.Map{
				"status":   "failed",
				"messages": "correct your input",
			})
	}

	errors := utilities.ValidateStruct(*params)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, err := h.ProductService.CreateProduct(ctx, params)
	if err != nil {
		return ctx.Status(500).
			JSON(fiber.Map{
				"status": "failed",
				"data":   nil,
			})
	}

	return ctx.Status(code).
		JSON(fiber.Map{
			"sesc":         "SUCCESS",
			"messageLocal": "Product disimpan",
		})
}

func NewProductHandler(service interfaces.ProductServiceInterface) Product {
	return &ProductHandler{
		ProductService: service,
	}
}
