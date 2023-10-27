package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/controller"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("/users", controller.AddUserController)
	e.GET("/users", controller.GetUserController)
	e.POST("/lokers", controller.AddLokerController)
	e.GET("/lokers", controller.GetLokerController)
	return e
}