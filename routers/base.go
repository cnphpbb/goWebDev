// base
package routers

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func Home(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"Code": 0, "Msg": "Request Succeed", "Data": APP_VER})
}

func NotFound() {
	return c.JSON(http.StatusNotFound, map[string]interface{}{"Code": 404, "Msg": "Request NotFound", "Data": ""})
}
