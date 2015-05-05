package main

import (
	"net/http"

	"github.com/cnphpbb/go_utils/modules/webfw"
	mw "github.com/cnphpbb/go_utils/modules/webfw/middleware"

	"github.com/cnphpbb/go_utils/routers"
)

const APP_VER = "0.0.1.0502 Beta"

func home(c *webfw.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"Code": 0, "Msg": "Request Succeed", "Data": APP_VER})
}

func main() {
	w := webfw.New()
	w.Use(mw.Logger)
	// Routes
	w.Get("/", routers.Home)
	w.NotFound(routers.NotFound)

	w.Run(":8989")
}
