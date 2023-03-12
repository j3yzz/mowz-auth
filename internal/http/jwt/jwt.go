package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/j3yzz/mowz/internal/model"
	"github.com/j3yzz/mowz/internal/store/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	AccessTokenSecret string
}

type JWT struct {
	Config
	UserStore user.User
}

type Claim struct {
	jwt.StandardClaims
}

func (j JWT) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
			userWithToken, _ := j.RetrieveByToken(tokenString, c)
			ctx := context.WithValue(c.Request().Context(), "user", userWithToken)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func (j JWT) RetrieveByToken(signedToken string, c echo.Context) (*model.UserWithId, error) {
	id, err := j.ValidateToken(signedToken)
	if err != nil {
		return nil, c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	user, err := j.UserStore.FindById(id)
	if err != nil {
		return nil, c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return &user, nil
}

func (j JWT) NewAccessToken(u model.UserWithId) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "user",
		ExpiresAt: time.Now().Add(24 * time.Hour * 30 * 6).Unix(),
		Id:        uuid.New().String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "mowz-auth",
		NotBefore: time.Now().Unix(),
		Subject:   fmt.Sprint(u.Id),
	})

	encodedToken, err := token.SignedString([]byte(j.AccessTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign a token: %w", err)
	}

	return encodedToken, nil
}

func (j JWT) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Config.AccessTokenSecret), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return "", errors.New("token expired")
	}

	return claims.Subject, nil
}
