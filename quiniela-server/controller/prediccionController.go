package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func GetPrediccion(c *fiber.Ctx) error {
	// idEvento
	var id string = c.Params("id")
	var prediccion models.PREDICCION

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unaunthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	query := fmt.Sprintf("SELECT * FROM PREDICCION WHERE idEvento = %s AND idUsuario = %s", id, claims.Issuer)

	rows,err := database.ExecuteQuery(query)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}
	for rows.Next() {
		rows.Scan(&prediccion.IdPrediccion, &prediccion.Local, &prediccion.Visitante,
			&prediccion.IdEvento, &prediccion.IdUsuario)
	}
	defer rows.Close()

	return c.JSON(prediccion)
}

func PostPrediccion(c *fiber.Ctx) error {
	var data models.PREDICCION_USUARIO

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	fmt.Println("SIIIIIIIIIIIIIIIIIIaaI")


	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unaunthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	query := fmt.Sprintf("CALL sp_insert_prediccion(%s,%s,%s,%s)",
		data.Local, data.Visitante, data.IdEvento, claims.Issuer)

	fmt.Println("****************PREDICCION: ", query)

	err = database.Execute(query)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error":   err,
		})
	}

	return c.JSON(data)
}

//zSf8rXfarK
//aboxerz@tiny.cc