package homepage

import (
	"github.com/labstack/echo"

	"akawork.io/module/homepage/controller"
)

var mHomePageController *controller.HomePageController

/**
 * Initializes module
 */
func Initialize(e *echo.Echo) {
	mHomePageController = &controller.HomePageController{}

	// New router
	InitRoute(e)
}

func InitRoute(e *echo.Echo) {
	gv2 := e.Group("arrow/api/v1.0/")

	// Api for Account
	gv2.GET("greeting", mHomePageController.Greeting)
	gv2.GET("user", mHomePageController.User)
	gv2.GET("parent", mHomePageController.Parent)

	e.GET("healthcheck", mHomePageController.GetHealthcheck)
}
