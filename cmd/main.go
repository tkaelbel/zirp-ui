package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/tkaelbel/zirp-ui/handler"
)

var store *session.Store

const (
	AUTH_KEY       string = "authenticated"
	USER_ID        string = "user_id"
	FROM_PROTECTED string = "from_protected"
	TZONE_KEY      string = "time_zone"
)

func main() {
	app := fiber.New()

	// app.Use(middleware.NotFoundMiddleware)
	// TODO: probably configure this
	// app.Use(cors.New())
	app.Use(logger.New())

	// Provide a minimal config
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./static"),
	}))

	handler.SetupRoutes(app)

	/* ↓ Not Found Management - Fallback Page ↓ */
	// app.Get("/*", flagsMiddleware, func(c *fiber.Ctx) error {

	// 	return fiber.NewError(
	// 		fiber.StatusNotFound,
	// 		"error 404: not found",
	// 	)
	// })

	log.Fatal(app.Listen(":3000"))
}

// flagsMiddleware is a middleware for handling different behaviors
// of non protected pages, specifically not allowing an already
// logged in user to log in or register again.
func flagsMiddleware(c *fiber.Ctx) error {
	sess, _ := store.Get(c)
	userId := sess.Get(USER_ID)
	if userId == nil {
		c.Locals(FROM_PROTECTED, false)

		return c.Next()
	}

	c.Locals(FROM_PROTECTED, true)

	return c.Next()
}
