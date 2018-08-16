package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()
	//Criação de rotas
	App.GET("/", home)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
