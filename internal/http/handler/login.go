package handler

import (
	"fmt"
	"github.com/j3yzz/mowz/internal/http/jwt"
	"github.com/j3yzz/mowz/internal/http/request"
	"github.com/j3yzz/mowz/internal/model"
	"github.com/j3yzz/mowz/internal/store/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Login struct {
	Store user.User
	JWT   jwt.JWT
}

func (l Login) Handle(c echo.Context) error {
	var rq request.Login
	var u model.UserWithId

	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := rq.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u, err := l.Store.FindByEmail(rq.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if u.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("email %s does not exist", rq.Email))
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(rq.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "incorrect password")
	}

	t, err := l.JWT.NewAccessToken(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"jwt":     t,
	})
}

func (l Login) Register(g *echo.Group) {
	g.POST("/login", l.Handle)
}
