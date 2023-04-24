package middlewares

import (
	"github.com/j3yzz/mowz/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

const USER_ROLE = "user"

func CheckUserHasRoleMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUser := c.Request().Context().Value("user").(*model.UserWithRole)
			if currentUser.Role != USER_ROLE {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"message": "user_does_not_have_access",
				})
			}

			return next(c)
		}
	}
}
