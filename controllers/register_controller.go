package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Register struct {
	ID         uint64 `json:"id"`
	IDCategory uint64 `json:"id_category"`
	Details    string `json:"details"`
}

type Registrys []Register

var registerList Registrys

func GetRegistrys(c echo.Context) error {
	return c.JSON(http.StatusOK, registerList)
}

func CreateRegister(c echo.Context) error {
	register := Register{}
	err := c.Bind(&register)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	registerList = append(registerList, register)
	return c.JSON(http.StatusCreated, registerList)
}

func DeleteRegister(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.String(http.StatusNoContent, string(id))
}
