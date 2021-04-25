package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func GetDeportes(c *fiber.Ctx) error {
	var deportes = models.ArrayDeporte{}

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
	var deporte models.DEPORTE

	rows,err := database.ExecuteQuery("SELECT * FROM DEPORTE")

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&deporte.IdDeporte, &deporte.Nombre, &deporte.Imagen, &deporte.Color)
		deportes = append(deportes, deporte)
	}
	defer rows.Close()

	return c.JSON(deportes)
}

func PostDeporte(c *fiber.Ctx) error {
	var data models.DEPORTE
	if err:= c.BodyParser(&data); err != nil {
		return err
	}
	query := fmt.Sprintf("CALL sp_insert_deporte('%s','%s','%s')", data.Nombre, data.Imagen, data.Color)


	_, err := database.ExecuteQuery(query)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	return c.JSON(data)
}

func PutDeporte(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var data models.DEPORTE
	if err:= c.BodyParser(&data); err != nil {
		return err
	}
	query := fmt.Sprintf("CALL sp_update_deporte('%s','%s',%s)", data.Imagen, data.Color, id)
	fmt.Println(query)
	_, err := database.ExecuteQuery(query)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update Successful",
	})
}

func DeleteDeporte(c *fiber.Ctx) error {
	var id string = c.Params("id")
	query := fmt.Sprintf("CALL sp_delete_deporte(%s)", id)
	_, err := database.ExecuteQuery(query)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Delete Successful",
	})
}

/*
	"idDeporte": 9
	"name": "prueba",
	"img": "xp",
	"color": "gris"

{
	"name": "prueba",
	"img": "xp",
	"color": "gris"
}
*/