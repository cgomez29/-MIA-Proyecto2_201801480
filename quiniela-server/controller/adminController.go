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
	var jornada models.Jornada
	var deporte models.DEPORTE
	var evento models.EVENTO

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

			//Obteniendo id Jornada actual
			query = fmt.Sprintf("SELECT idJornada FROM JORNADA WHERE name = '%s' AND idTemporada = %s",
				j.Jornada, strconv.Itoa(int(temporada.IdTemporada)))

			rows,err := database.ExecuteQuery(query)
			if err != nil  {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": "Error running query",
					"error": err,
				})
			}

			for rows.Next() {
				rows.Scan(&jornada.IdJornada)
			}

			for _, k:= range j.Predicciones {
				// INSERTANDO DEPORTES
				query = fmt.Sprintf("CALL sp_insert_deporte_bl('%s')", k.Deporte)
				database.ExecuteQuery(query)

				//Obteniendo id Deporte
				query = fmt.Sprintf("SELECT idDeporte FROM DEPORTE WHERE nombre = '%s'",
						k.Deporte)

				rows,err = database.ExecuteQuery(query)
				if err != nil  {
					c.Status(fiber.StatusInternalServerError)
					return c.JSON(fiber.Map{
						"message": "Error running query",
						"error": err,
					})
				}

				for rows.Next() {
					rows.Scan(&deporte.IdDeporte)
				}

				//CREANDO EVENTOS
				query = fmt.Sprintf("CALL sp_insert_evento_bl('%s','%s','%s',%s,%s)",
						k.Fecha, k.Local, k.Visitante, strconv.Itoa(jornada.IdJornada), strconv.Itoa(int(deporte.IdDeporte)))
				database.ExecuteQuery(query)

				//OBTENIENDO id de EVENTO
				query = fmt.Sprintf("SELECT idEvento FROM EVENTO WHERE idJornada = %s AND idDeporte = %s AND fecha_hora = TO_TIMESTAMP('%s', 'YYYY-MM-DD HH24:MI:SS.FF')",
					strconv.Itoa(jornada.IdJornada), strconv.Itoa(int(deporte.IdDeporte)), k.Fecha)

				rows,err = database.ExecuteQuery(query)
				if err != nil  {
					c.Status(fiber.StatusInternalServerError)
					return c.JSON(fiber.Map{
						"message": "Error running query",
						"error": err,
					})
				}

				for rows.Next() {
					rows.Scan(&evento.IdEvento)
				}

				//CREANDO RESULTADOS
				query = fmt.Sprintf("CALL sp_insert_resultado_bl(%s,%s,%s)",
					strconv.Itoa(int(k.Resultado.Visitante)), strconv.Itoa(int(k.Resultado.Local)),
						strconv.Itoa(evento.IdEvento))
				database.ExecuteQuery(query)

				//CREANDO PREDICCIONES
				query = fmt.Sprintf("CALL sp_insert_prediccion_bl(%s,%s,%s,%s)",
					strconv.Itoa(int(k.Prediccion.Local)), strconv.Itoa(int(k.Prediccion.Visitante)),
					strconv.Itoa(evento.IdEvento), strconv.Itoa(int(user.IdUsuario)))
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