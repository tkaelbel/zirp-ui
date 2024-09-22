package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sujit-baniya/flash"
	"github.com/tkaelbel/zirp-ui/view"
)

func GetRegister(ctx *fiber.Ctx) error {

	fromProtected := ctx.Locals(FROM_PROTECTED).(bool)

	rindex := view.RegisterIndex(fromProtected)
	register := view.Register("/register", fromProtected, false, flash.Get(ctx), rindex)

	handler := adaptor.HTTPHandler(templ.Handler(register))

	return handler(ctx)
}
