package route

import (
	"../controller"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	//routes
	app.Get("/", controller.Welcome)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
	app.Get("/api/img/:id", controller.ViewImg)

	// DEPORTES
	app.Get("/api/deporte", controller.GetDeportes)
	app.Post("/api/deporte", controller.PostDeporte)
	app.Put("/api/deporte/:id", controller.PutDeporte)
	app.Delete("/api/deporte/:id", controller.DeleteDeporte)

	//CARGAR ARCHIVO
	app.Post("/api/bulkload",controller.BulkLoad)

}
