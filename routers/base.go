// base
package routers

import (
	"fmt"
)

func Home(c *webfw.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"Code": 0, "Msg": "Request Succeed", "Data": APP_VER})
}

func NotFound() {
	return c.JSON(http.StatusNotFound, map[string]interface{}{"Code": 404, "Msg": "Request NotFound", "Data": ""})
}
