package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"strconv"
	"time"

	//"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Welcome(c *fiber.Ctx) error {
	return c.SendString("QUINIELA - API  👋!")
}

func Register(c *fiber.Ctx) error {
	var data models.USUARIO
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//password, _:= bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	query := fmt.Sprintf("CALL sp_insert_usuario('%s','%s','%s','%s','%s','%s','%s')",
		data.Username, data.Password, data.Name, data.Surname, data.FechaNacimiento,
		data.Photo, data.Email)
	fmt.Println(query)


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

func Login(c *fiber.Ctx) error  {
	var data models.USUARIO
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	query := fmt.Sprintf("CALL sp_auth('%s','%s')", data.Username, data.Password)

	fmt.Println(query)

	_, err := database.ExecuteQuery(query)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect username or password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(data.IdRol)),
		Id:     strconv.Itoa(int(data.IdUsuario)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	var user models.USUARIO

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

	query := fmt.Sprintf("SELECT * FROM v_user WHERE idUsuario = %s;", claims.Id)

	rows,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&user.IdUsuario, &user.Username, &user.Name, &user.Surname,
					&user.FechaNacimiento, &user.Email, &user.Photo)
	}
	defer rows.Close()

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}