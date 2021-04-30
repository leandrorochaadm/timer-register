package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leandrorochaadm/time-register/config"
)

func Generate() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", nil)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
