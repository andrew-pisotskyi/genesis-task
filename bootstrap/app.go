package bootstrap

import (
	"github.com/andrew-pisotskyi/genesis-task/api/route"
	"github.com/andrew-pisotskyi/genesis-task/config"
	"github.com/labstack/echo/v4"
	"log"
)

type Application struct {
	config *config.Config
}

func App(c *config.Config) Application {
	app := &Application{}
	app.config = c
	return *app
}

func (app *Application) Run(e *echo.Echo) {
	router := route.NewRouter(app.config, e)
	router.Setup()

	log.Fatal(e.Start(app.config.ServerAddress))
}
