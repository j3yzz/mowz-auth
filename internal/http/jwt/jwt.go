package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/j3yzz/mowz/internal/http/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	AccessTokenSecret string
}

type JWT struct {
	Config
}

func (j JWT) Middleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ContextKey:    common.UserContextKey,
		SigningKey:    []byte(j.Config.AccessTokenSecret),
		SigningMethod: jwt.SigningMethodHS256.Name,
		Claims:        &jwt.StandardClaims{},
		TokenLookup:   "header:Authorization",
	})
}
