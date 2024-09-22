package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

const (
	AUTH_KEY       string = "authenticated"
	USER_ID        string = "user_id"
	FROM_PROTECTED string = "from_protected"
	TZONE_KEY      string = "time_zone"
)

func SetupRoutes(app *fiber.App) {
	/* Sessions Config */
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})

	app.Get("/home", flagsMiddleware, GetHome)
	app.Get("/login", flagsMiddleware, GetLogin)
	app.Get("/register", flagsMiddleware, GetRegister)

	// router := app.Group("/")
	// NewHomeHandler(router, nil)
	// handler.NewLoginHandler("/", nil)

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
