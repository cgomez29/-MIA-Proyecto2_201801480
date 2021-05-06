package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/gofiber/fiber"
)

func BulkLoad(c *fiber.Ctx) error  {
	var data models.Temporal
	var user models.USUARIO

	if err:= c.BodyParser(&data); err != nil {
		return err
	}
	var query string
	//INSERTANDO USUARIO
	query = fmt.Sprintf("CALL sp_insert_usuario_bl('%s','%s','%s','%s')",
		data.Username, data.Password, data.Nombre, data.Apellido)
	_,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Usuario repetido",
		})
	}

	// Recuperando id de usuario ingresado
	query = fmt.Sprintf("SELECT idUsuario FROM USUARIO WHERE username = %s", data.Username)

	rows,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&user.IdUsuario)
	}

	for _, i:= range data.Resultados {
		//INSERTANDO TEMPORADAS
		//formando la fecha
		dateTemporada := fmt.Sprintf("%s%s%s%s",i.Temporada[0:4],"/",i.Temporada[6:],"/01")
		query = fmt.Sprintf("CALL sp_insert_temporada_bl('%s','%s')", i.Temporada, dateTemporada)
		database.ExecuteQuery(query)
		for _, j:= range i.Jornadas {

			for _, k:= range j.Predicciones {
				// INSERTANDO DEPORTES
				query = fmt.Sprintf("CALL sp_insert_deporte_bl('%s')", k.Deporte)
				database.ExecuteQuery(query)
			}

		}
	}

	/*

	//*/

	return c.JSON(fiber.Map{
		"message": "Successful",
	})
}