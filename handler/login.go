package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sujit-baniya/flash"
	"github.com/tkaelbel/zirp-ui/view"
)

func GetLogin(ctx *fiber.Ctx) error {
	fromProtected := ctx.Locals(FROM_PROTECTED).(bool)

	lindex := view.LoginIndex(fromProtected)
	login := view.Login(
		" | Login", fromProtected, false, flash.Get(ctx), lindex,
	)

	handler := adaptor.HTTPHandler(templ.Handler(login))

	// if ctx.Method() == "POST" {
	// 	// obtaining the time zone from the POST request of the login form
	// 	tzone := ""
	// 	if len(ctx.GetReqHeaders()["X-Timezone"]) != 0 {
	// 		tzone = ctx.GetReqHeaders()["X-Timezone"][0]
	// 		// fmt.Println("Tzone:", tzone)
	// 	}

	// 	var (
	// 		user models.User
	// 		err  error
	// 	)
	// 	fm := fiber.Map{
	// 		"type": "error",
	// 	}

	// 	// notice: in production you should not inform the user
	// 	// with detailed messages about login failures
	// 	if user, err = models.CheckEmail(c.FormValue("email")); err != nil {
	// 		// fmt.Println(err)
	// 		if strings.Contains(err.Error(), "no such table") ||
	// 			strings.Contains(err.Error(), "database is locked") {
	// 			// "no such table" is the error that SQLite3 produces
	// 			// when some table does not exist, and we have only
	// 			// used it as an example of the errors that can be caught.
	// 			// Here you can add the errors that you are interested
	// 			// in throwing as `500` codes.
	// 			return fiber.NewError(
	// 				fiber.StatusServiceUnavailable,
	// 				"database temporarily out of service",
	// 			)
	// 		}
	// 		fm["message"] = "There is no user with that email"

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	err = bcrypt.CompareHashAndPassword(
	// 		[]byte(user.Password),
	// 		[]byte(c.FormValue("password")),
	// 	)
	// 	if err != nil {
	// 		fm["message"] = "Incorrect password"

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	session, err := store.Get(c)
	// 	if err != nil {
	// 		fm["message"] = fmt.Sprintf("something went wrong: %s", err)

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	session.Set(AUTH_KEY, true)
	// 	session.Set(USER_ID, user.ID)
	// 	session.Set(TZONE_KEY, tzone)

	// 	err = session.Save()
	// 	if err != nil {
	// 		fm["message"] = fmt.Sprintf("something went wrong: %s", err)

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	fm = fiber.Map{
	// 		"type":    "success",
	// 		"message": "You have successfully logged in!!",
	// 	}

	// 	return flash.WithSuccess(c, fm).Redirect("/todo/list")
	// }

	return handler(ctx)

}
