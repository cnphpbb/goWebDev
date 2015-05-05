package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/cnphpbb/go_utils/routers"
)

const APP_VER = "0.0.1.0502 Beta"

func main() {
	w := webfw.New()
	w.Use(mw.Logger)
	// Routes
	w.Get("/", routers.Home)
	w.NotFound(routers.NotFound)

	w.Run(":8989")
}
