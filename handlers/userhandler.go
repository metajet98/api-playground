package handlers

import (
	"api-playground/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "DKM sai param")
	}

	err := models.InsertNewUser(u)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, "Lỗi khi tạo user")
	}
	return c.JSON(http.StatusOK, u)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetUsers())
}
