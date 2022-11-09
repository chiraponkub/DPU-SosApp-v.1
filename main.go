package main

import (
	"fmt"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
)

func main() {

	ctrl := restapi.NewController()

	err := ctrl.LoadConfigFile()
	if err != nil {
		panic("LoadConfigFile from yml file error: " + err.Error())
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(200)))

	e.GET("/", ctrl.Test)

	r := e.Group(config.GetString("service.endpoint"))
	r.GET("/", func(c echo.Context) error {
		fmt.Println(config.GetString("service.endpoint"))
		return c.JSON(200, "Ok")
	})

	// "/user"
	u := r.Group(config.GetString("role.user"))
	{
		u.GET("/name", func(c echo.Context) error {
			return c.JSON(200, "Name")
		})
	}

	a := r.Group(config.GetString("role.admin"))
	{
		a.GET("/role", ctrl.GetRoleList)
		a.POST("/role", ctrl.AddRole)
		//a.PUT("/role", AddRole)
		//a.DELETE("/role", AddRole)
	}

	e.Logger.Fatal(e.Start(":80"))
}
