package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/gofiber/fiber"
	"strconv"
)

func BulkLoad(c *fiber.Ctx) error  {
	var data models.Temporal
	var user models.USUARIO
	var temporada models.Temporada

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
	query = fmt.Sprintf("SELECT idUsuario FROM USUARIO WHERE username = '%s'", data.Username)

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

		//ASIGNANDO  tier para temporada actual
		//Obteniendo id temporada

		query = fmt.Sprintf("SELECT idTemporada FROM TEMPORADA WHERE nombre = '%s'", i.Temporada )

		rows,err := database.ExecuteQuery(query)
		if err != nil  {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Error running query",
				"error": err,
			})
		}

		for rows.Next() {
			rows.Scan(&temporada.IdTemporada)
		}
		var idMembresia int

		if i.Tier == "gold" {
			idMembresia = 1
		} else if i.Tier == "silver" {
			idMembresia = 2
		} else {
			idMembresia = 3
		}

		query = fmt.Sprintf("CALL sp_insert_detalle_usuario_bl(%s,%s,%s)",
			strconv.Itoa(int(temporada.IdTemporada)), strconv.Itoa(idMembresia), strconv.Itoa(int(user.IdUsuario)))
		database.ExecuteQuery(query)

		for _, j:= range i.Jornadas {
			//CREANDO JORNADAS
			semana := fmt.Sprintf("%s",i.Temporada[6:])
			query = fmt.Sprintf("CALL sp_insert_jornada_bl('%s','%s',%s,%s)",
				j.Jornada, dateTemporada, semana, strconv.Itoa(int(temporada.IdTemporada)))
			database.ExecuteQuery(query)

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