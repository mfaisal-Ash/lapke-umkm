package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidiq200/lapke-umkm/controller/helper/json"
)

func ErrHandler(ctx *fiber.Ctx, err error) error {
	resp := json.ReturnData{
		fiber.StatusInternalServerError,
		false,
		"mengalami Error " + err.Error(),
		nil,
	}
	if e, ok := err.(*fiber.Error); ok {
		resp.Code = e.Code
		resp.Status = e.Message
	}
	return resp.WriteToBody(ctx)
}
