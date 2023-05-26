package route

import (
	"github.com/andrew-pisotskyi/genesis-task/config"

	"github.com/labstack/echo/v4"
)

type Router struct {
	config *config.Config
	echo   *echo.Echo
}

func NewRouter(config *config.Config, e *echo.Echo) *Router {
	return &Router{
		config,
		e,
	}
}

func (r *Router) Setup() {
	publicRouter := r.echo.Group("/api")
	r.NewBtcRateRoute(publicRouter)
}
