package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"myapp/database"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	// path ---

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserById)

	// --
	println("http://localhost:1323")
	e.Logger.Fatal(e.Start(":1323"))
}

func getUsers(c echo.Context) error {

	db := database.Connect()

	var users []User

	result := db.Find(&users)

	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func getUserById(c echo.Context) error {

	db := database.Connect()

	var user User

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	result := db.First(&user, id)

	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, user)
}