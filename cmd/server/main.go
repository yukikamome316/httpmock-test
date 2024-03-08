package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yukikamome316/httpmock-test/internal/client"
	"github.com/yukikamome316/httpmock-test/internal/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	netClient := client.NewClient("https://example.com", 300)
	h := handler.NewHandler(netClient)

	e.GET("/posts/:id", h.GetPosts)

	e.Logger.Fatal(e.Start(":8080"))
}
