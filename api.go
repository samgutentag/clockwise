package main

import (
	"github.com/alecsavvy/clockwise/db"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/labstack/echo/v4"
)

type ApiServer struct {
	rpc     *local.Local
	queries *db.Queries
	node    string
}

func (s *ApiServer) Start() error {
	e := echo.New()

	e.HideBanner = true

	e.GET("/health_check", s.getHealth)

	// e.GET("/users")
	// e.GET("/tracks")

	// e.POST("/users/:handle")
	// e.POST("/users/:handle/tracks/:title")

	return e.Start(":8080")
}

func (s *ApiServer) getHealth(c echo.Context) error {
	res := make(map[string]string)

	res["node_id"] = s.node
	res["status"] = "up"

	return c.JSON(200, res)
}
