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

func GetDeporte(c *fiber.Ctx) error {
	var id string = c.Params("id")

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

	query := fmt.Sprintf("SELECT * FROM DEPORTE WHERE idDeporte = %s" , id)

	rows,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&deporte.IdDeporte, &deporte.Nombre, &deporte.Imagen, &deporte.Color)
	}
	defer rows.Close()
	return c.JSON(deporte)
}

func PostDeporte(c *fiber.Ctx) error {
	var data models.DEPORTE

	if err := c.BodyParser(&data); err != nil {
		return err
	}

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

	// save photo
	var nameImg string

	if form, err := c.MultipartForm(); err == nil {
		// Get all files from "documents" key:
		files := form.File["file"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			path := fmt.Sprintf("./files/%s_%s", data.Nombre, file.Filename)
			nameImg = fmt.Sprintf("%s_%s", data.Nombre, file.Filename)
			// Save the files to disk:
			if err := c.SaveFile(file,path); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	// UPDATE
	query := fmt.Sprintf("CALL sp_insert_deporte('%s','%s','%s')", data.Nombre, nameImg, data.Color)
	err = database.Execute(query)

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

	if err := c.BodyParser(&data); err != nil {
		return err
	}

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

	// save photo
	var nameImg string

	nameImg = data.Imagen

	if form, err := c.MultipartForm(); err == nil {
		// Get all files from "documents" key:
		files := form.File["file"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			path := fmt.Sprintf("./files/%s_%s", data.Nombre, file.Filename)
			nameImg = fmt.Sprintf("%s_%s", data.Nombre, file.Filename)
			// Save the files to disk:
			if err := c.SaveFile(file,path); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	// UPDATE
	query := fmt.Sprintf("CALL sp_update_deporte('%s','%s',%s)",nameImg, data.Color, id)
	err = database.Execute(query)

	fmt.Println("UPDATE",query)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}
	return c.JSON(data)
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