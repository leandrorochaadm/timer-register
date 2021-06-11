package registrys

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/leandrorochaadm/time-register/database"
)

type Register struct {
	ID         uint64    `json:"id"`
	Datetime   time.Time `json:"datetime"`
	IDCategory uint64    `json:"id_category"`
	Details    string    `json:"details"`
	Duration   int64     `json:"duration"`
	Time       string    `json:"time"`
}

type Registrys []Register

var ciclePrimary bool

func GetRegistrys(c echo.Context) error {
	var registerList Registrys
	rows, err := database.DBClient.
		Query("SELECT id, datetime, details, id_category FROM registrys;")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			"Não foi possível listar os registros \n"+err.Error())
	}
	var datetimePrevious time.Time
	for rows.Next() {
		var singleRegister Register
		if err := rows.Scan(&singleRegister.ID, &singleRegister.Datetime, &singleRegister.Details,
			&singleRegister.IDCategory); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity,
				"Não foi possível deserializer os registros \n"+err.Error())

		}

		if ciclePrimary {
			singleRegister.Duration = int64(singleRegister.Datetime.Sub(datetimePrevious) / time.
				Second)
			singleRegister.Time = singleRegister.Datetime.Sub(datetimePrevious).String()
		} else {
			singleRegister.Duration = int64(singleRegister.Datetime.Sub(singleRegister.Datetime) / time.Second)
			singleRegister.Time = singleRegister.Datetime.Sub(singleRegister.Datetime).String()
		}

		datetimePrevious = singleRegister.Datetime
		// singleRegister.Duration = (25*time.Hour+ 3*time.Second).String()
		registerList = append(registerList, singleRegister)
		ciclePrimary = true
	}
	return c.JSON(http.StatusOK, registerList)
	// return c.JSON(http.StatusOK, registerList)
}

func CreateRegister(c echo.Context) error {
	register := Register{}
	err := c.Bind(&register)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	res, err := database.DBClient.Exec("INSERT INTO registrys (id_category, details) VALUES (?,?);",
		register.IDCategory, register.Details)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			"Não foi possível gravar esse registro: \n"+err.Error())
	}
	id, _ := res.LastInsertId()

	// registerList = append(registerList, register)
	return c.JSON(http.StatusCreated, id)
}

func DeleteRegister(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusUnprocessableEntity,
			"É necessário informar a tarefa que deseja deletar\n "+err.Error())
	}
	fmt.Print(id)

	_, err = database.DBClient.Exec("delete from registrys where id = ?", id)
	if err != nil {
		echo.NewHTTPError(http.StatusUnprocessableEntity,
			"Não foi possível deletar essa tarefa \n "+err.Error())
	}
	return c.String(http.StatusNoContent, "")
}
