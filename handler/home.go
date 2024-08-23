package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sujit-baniya/flash"
	"github.com/tkaelbel/zirp-ui/view"
)

func GetHome(ctx *fiber.Ctx) error {

	fromProtected := ctx.Locals(FROM_PROTECTED).(bool)

	hindex := view.HomeIndex(fromProtected)
	home := view.Home("/home", fromProtected, false, flash.Get(ctx), hindex)

	handler := adaptor.HTTPHandler(templ.Handler(home))

	return handler(ctx)
	// name := "World"
	// return Render(ctx, view.Home(name))
}
