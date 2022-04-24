package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"myapp/database"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	// SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// 	SigningKey: []byte("secret"),
	// }))

	e.Pre(middleware.RemoveTrailingSlash())
	
	// --------------------------------------------------------------------------------
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	

	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserById)
	e.DELETE("/users/:id", deleteUser)
	// --------------------------------------------------------------------------------

	// export routes data to json file
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)
	
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

func deleteUser(c echo.Context) error {
	
	db := database.Connect()

	var user User

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	result := db.Delete(&user, id)

	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, user)
}
