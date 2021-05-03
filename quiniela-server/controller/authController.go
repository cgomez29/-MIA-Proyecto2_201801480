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

func Login(c *fiber.Ctx) error  {
	var data models.USUARIO
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//query := fmt.Sprintf("CALL sp_auth('%s','%s')", data.Username, data.Password)
	query := fmt.Sprintf("SELECT idUsuario, idRol FROM USUARIO WHERE username = '%s' AND password = '%s'", data.Username, data.Password)

	rows, err := database.ExecuteQuery(query)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect username or password",
		})
	}

	for rows.Next() {
		fmt.Println(rows)
		rows.Scan(&data.IdUsuario, &data.IdRol)
	}
	defer rows.Close()

	//fmt.Println("result: " ,data)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(data.IdUsuario)),
		Id: strconv.Itoa(int(data.IdRol)),
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
		"idRol": data.IdRol,
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

	query := fmt.Sprintf("SELECT * FROM v_user WHERE idUsuario = %s", claims.Issuer)

	rows,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&user.IdUsuario, &user.Username, &user.Name, &user.Surname, &user.Tier,
					&user.FechaNacimiento, &user.Email, &user.Photo, &user.IdRol)
	}
	defer rows.Close()
	fmt.Println(user)
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

func Register(c *fiber.Ctx) error {
	// register
	var data models.USUARIO
	if err := c.BodyParser(&data); err != nil {
		return err
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
	query := fmt.Sprintf("CALL sp_insert_usuario('%s','%s','%s','%s','%s','%s','%s')",
		data.Username, data.Password, data.Name, data.Surname, data.FechaNacimiento,
		nameImg, data.Email)
	fmt.Println("REGISTER: ",query)

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

func ViewImg(c *fiber.Ctx) error {
	var name string = c.Params("id")
	img := fmt.Sprintf("./files/%s", name)
	return c.SendFile(img)
}