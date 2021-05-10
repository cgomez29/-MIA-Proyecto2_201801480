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

	//USUARIO
	app.Put("/api/user", controller.PutUser)
	app.Put("/api/user/membresia", controller.PutUserMembresia)

	// DEPORTES
	app.Get("/api/deporte", controller.GetDeportes)
	app.Get("/api/deporte/:id", controller.GetDeporte)
	app.Post("/api/deporte", controller.PostDeporte)
	app.Put("/api/deporte/:id", controller.PutDeporte)
	app.Delete("/api/deporte/:id", controller.DeleteDeporte)

	//CARGAR ARCHIVO
	app.Post("/api/bulkload",controller.BulkLoad)

}
