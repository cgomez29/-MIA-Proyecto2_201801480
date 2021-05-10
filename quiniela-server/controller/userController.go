package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

// PutUser Update user
func PutUser(c *fiber.Ctx) error {
	var data models.USUARIO
	if err := c.BodyParser(&data); err != nil {
		return err
	}

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

	// save photo
	var nameImg string

	nameImg = data.Photo

	if form, err := c.MultipartForm(); err == nil {
		// Get all files from "documents" key:
		files := form.File["file"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			path := fmt.Sprintf("./files/%s_%s", data.Username, file.Filename)
			nameImg = fmt.Sprintf("%s_%s", data.Username, file.Filename)
			// Save the files to disk:
			if err := c.SaveFile(file,path); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	// register
	//password, _:= bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	query := fmt.Sprintf("CALL sp_update_usuario('%s','%s','%s','%s','%s','%s','%s',%s)",
		data.Username, data.Password, data.Name, data.Surname, data.Fecha,
		nameImg, data.Email, claims.Issuer)

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

// PutUserMembresia Update tier the user
func PutUserMembresia(c *fiber.Ctx) error {
	var data models.USUARIO
	if err := c.BodyParser(&data); err != nil {
		return err
	}

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

	query := fmt.Sprintf("CALL sp_update_usuario_membresia('%s',%s)",
		data.Tier, claims.Issuer)
	fmt.Println(query)
	err = database.Execute(query)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Successful",
	})
}
