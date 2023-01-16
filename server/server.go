package server

import (
	"servertestgo/handlers"

	"github.com/labstack/echo/v4"
)

func InitServer(PORT string) {
	//para crear el servidor http, usamos la librería echo por plasticidad.
	e := echo.New()
	//registramos las rutas
	e.GET("/", handlers.Home)
	e.GET("/estimate/:name", handlers.Estimate)
	e.POST("/swap/", handlers.Swap)
	e.GET("/swap/", handlers.SwapAll)
	e.Logger.Fatal(e.Start(PORT))
}
