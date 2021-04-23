package route

import (
	"../controller"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	//routes
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
}
