package controller

import (
	"../database"
	"../models"
	"../models2"
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
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//password, _:= bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.ADMIN{
		Username: data["username"],
		Password: data["password"],
		Email:	  data["email"],
		//Password: password
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error  {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models2.ADMIN

	database.DB.Where("email=?",data["email"]).First(&user)

	if user.Idadmin == 10 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if user.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Idadmin)),
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

	var user models2.ADMIN

	database.DB.Where("idadmin = ?",claims.Issuer).First(&user)

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