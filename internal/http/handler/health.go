package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
}

func (h Health) Handle(c echo.Context) error {
	return c.String(http.StatusOK, "everything is fine.")
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health", h.Handle)
}
