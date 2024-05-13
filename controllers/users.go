package services

import (
	"errors"
	"lecho/database"
	"net/http"
	"strconv"

	m "lecho/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context) error {

	db := database.Connect()

	var users []m.User

	result := db.Find(&users)

	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {

	db := database.Connect()

	var user m.User

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	result := db.First(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// return c.String(http.StatusNotFound, "User not found")
			user = m.User{
				Id:    id,
				Name:  "John Doe",
				Email: "john@example.com",
			}
			db.Create(&user)

			return c.JSON(http.StatusOK, user)
		} else {
			return c.String(http.StatusInternalServerError, result.Error.Error())
		}
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {

	db := database.Connect()

	var user m.User

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
