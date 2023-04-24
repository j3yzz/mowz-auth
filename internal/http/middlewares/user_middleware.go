package middlewares

import (
	"github.com/j3yzz/mowz/internal/model"
	"github.com/j3yzz/mowz/internal/store/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const USER_ROLE = "user"

func CheckUserHasRoleMiddleware(db *user.MysqlUser) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUser := c.Request().Context().Value("userWithRole").(*model.UserWithId)
			userWithRole, err := db.FindUserWithRole(strconv.Itoa(currentUser.Id))
			if err != nil {
				return err
			}

			if userWithRole.Role != USER_ROLE {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"message": "user_does_not_have_access",
				})
			}

			return next(c)
		}
	}
}
