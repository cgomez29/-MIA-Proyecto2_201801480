package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func GetEventos(c *fiber.Ctx) error {
	var eventos = models.ArrayEventCalendar{}

	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unaunthenticated",
		})
	}

	var evento models.EVENTOCALENDAR
	var eventoCalendar models.EVENTCALENDAR

	rows,err := database.ExecuteQuery("SELECT * FROM EVENTO")

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&evento.IdEvento, &evento.FechaHora, &evento.Estado, &evento.Local,
			&evento.Visitante, &evento.Color, &evento.IdJornada, &evento.IdDeporte)
		//2019-06-06T19:16:00Z

		dateEvent := fmt.Sprintf("%s %s",evento.FechaHora[0:10], evento.FechaHora[11:16])
		title := fmt.Sprintf("%s VS %s",evento.Local, evento.Visitante)

		eventoCalendar.IdEvento = evento.IdEvento
		eventoCalendar.Title = title
		eventoCalendar.FechaHora = dateEvent
		eventoCalendar.Color = evento.Color

		eventos = append(eventos, eventoCalendar)
	}
	defer rows.Close()

	return c.JSON(eventos)
}