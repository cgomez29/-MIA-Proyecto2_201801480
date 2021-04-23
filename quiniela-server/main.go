package main

import (
	"./database"
	"./route"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	route.Setup(app)

	app.Listen(":8000")
}


/*func main() {
	database.Connect()

	router := mux.NewRouter().StrictSlash(true)

	route.Routes(router)
	//start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}*/

