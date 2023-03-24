package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/param"
)

type ProductServiceInterface interface {
	CreateProduct(ctx *fiber.Ctx, params *param.Request) (int, error)
}
