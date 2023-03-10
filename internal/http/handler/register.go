package handler

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/j3yzz/mowz/internal/http/request"
	"github.com/j3yzz/mowz/internal/model"
	"github.com/j3yzz/mowz/internal/store/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Register struct {
	Store user.User
}

func (r Register) Handle(c echo.Context) error {
	var rq request.Register
	var mysqlErr *mysql.MySQLError

	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := rq.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rq.Password), 8)
	if err != nil {
		panic(err)
	}
	u := model.User{
		Name:     rq.Name,
		Email:    rq.Email,
		Password: string(hashedPassword),
	}

	if err := r.Store.Set(u); err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == uint16(user.ErrEmailDuplicateCode) {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("email %s already exists", u.Email))
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"email": u.Email,
			"name":  u.Name,
		},
	})
}

func (r Register) Register(g *echo.Group) {
	g.POST("/register", r.Handle)
}
