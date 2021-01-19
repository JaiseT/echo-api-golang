package main

import (
	"net/http"

	"aot-technologies.com/saltsecurities/echo/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Root route => handler
	e.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.POST("/echo", func(c echo.Context) error {
		//return controllers.PostHTTPRequest()
		req := new(models.HTTPRequest)
		//m := echo.Map{}
		if err := c.Bind(&req); err != nil {
			return err
		}
		return c.JSON(200, req)
	})
	// Run Server
	e.Logger.Fatal(e.Start(":3000"))
}
