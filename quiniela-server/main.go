package main

import (
	"./controller"
	"./database"
	"./route"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	app := fiber.New()

	database.Connect()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	route.Setup(app)

	//Crea automaticamente la temporada si no existe
	controller.TemporadaActual()
	controller.Rewards()
	app.Listen(":8000")
}


/*func main() {
	database.Connect()

	router := mux.NewRouter().StrictSlash(true)

	route.Routes(router)
	//start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}*/

