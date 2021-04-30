package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leandrorochaadm/time-register/config"
	"github.com/leandrorochaadm/time-register/controllers"
)

func Generate() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bem vindo ao Time Register !")
	})
	e.GET("/registrys", controllers.GetRegistrys)
	e.POST("/register", controllers.CreateRegister)
	e.DELETE("/register", controllers.DeleteRegister)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
