package router

import (
	"fmt"
	"github.com/leandrorochaadm/time-register/categories"
	"github.com/leandrorochaadm/time-register/registrys"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leandrorochaadm/time-register/config"
)

func Generate() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bem vindo ao Time Register !")
	})
	e.GET("/registrys", registrys.GetRegistrys)
	e.POST("/register", registrys.CreateRegister)
	e.DELETE("/register/:id", registrys.DeleteRegister)
	e.GET("/categories", categories.GetCategories)
	e.POST("/category", categories.CreateCategory)
	e.DELETE("/category/:id", categories.DeleteCategory)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
