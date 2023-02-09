package main

import (
	app "github.com/Sergei39/chat-server/internal/app/check_spam"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	impl := app.NewService()
	route(e, impl)

	e.Logger.Fatal(e.Start(":7000"))
}

func route(e *echo.Echo, service *app.Implementation) {
	e.POST("/get-permissions", service.GetPermissions)
}
