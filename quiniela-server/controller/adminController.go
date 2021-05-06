package controller

import (
	"../models"
	"fmt"
	"github.com/gofiber/fiber"
)

func BulkLoad(c *fiber.Ctx) error  {
	var data models.Temporal
	if err:= c.BodyParser(&data); err != nil {
		return err
	}
	//query := fmt.Sprintf("CALL sp_insert_deporte('%s','%s','%s')", data.Nombre, data.Imagen, data.Color)
	//_, err := database.ExecuteQuery(query)

	fmt.Println("NOMBRE ",data.Nombre)

	/*if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}*/

	return c.JSON(data)
}