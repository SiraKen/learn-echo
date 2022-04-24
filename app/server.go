package main

import (
	"net/http"
	"os"
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
	
	// --------------------------------------------------------------------------------
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserById)
	// --------------------------------------------------------------------------------
	
	// set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	// start server
	println("http://localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
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