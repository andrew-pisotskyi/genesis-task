package main

import (
	"github.com/andrew-pisotskyi/genesis-task/bootstrap"
	"github.com/andrew-pisotskyi/genesis-task/config"
	"github.com/labstack/echo/v4"
)

func main() {
	c := config.NewConfig()
	app := bootstrap.App(c)
	e := echo.New()

	app.Run(e)
}
