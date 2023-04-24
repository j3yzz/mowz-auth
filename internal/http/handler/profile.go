package handler

import (
	"github.com/j3yzz/mowz/internal/http/jwt"
	"github.com/j3yzz/mowz/internal/model"
	"github.com/j3yzz/mowz/internal/store/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Profile struct {
	Store user.User
	JWT   jwt.JWT
}

func (p Profile) Handle(c echo.Context) error {
	currentUser := c.Request().Context().Value("user").(*model.UserWithRole)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    currentUser,
	})
}

func (p Profile) Register(g *echo.Group) {
	g.GET("/profile", p.Handle)
}
