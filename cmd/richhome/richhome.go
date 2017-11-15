package main

import (
    "net/http"

    "github.com/microcats/richhome/config"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: config.LogFormat,
    }))

    e.GET("/", func(c echo.Context) error {
        e.Logger.Info("aaa")
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Info("aaa")
    e.Logger.Fatal(e.Start(":1323"))
}
