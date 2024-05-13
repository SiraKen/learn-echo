package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	c "lecho/controllers"
)

// Export routes data to json file
func exportRoutesToJson(r []*echo.Route) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	rUsers := e.Group("/users")
	rUsers.GET("/", c.GetUsers)
	rUsers.GET("/:id", c.GetUserById)
	rUsers.DELETE("/:id", c.DeleteUser)

	exportRoutesToJson(e.Routes())

	// set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	// start server
	println("http://localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
