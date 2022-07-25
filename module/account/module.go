package account

import (
	"time"

	"github.com/labstack/echo"

	"akawork.io/infrastructure/cache"
	"akawork.io/module/account/controller"
	"akawork.io/module/account/service"
)

var mAccountController *controller.AccountController

/**
 * Initializes module
 */
func Initialize(e *echo.Echo, cache cache.CacheManager, timeout time.Duration) {
	cacheService := service.NewCacheService(cache, timeout)
	mAccountController = &controller.AccountController{
		Cache:        cache,
		CacheService: cacheService,
	}

	// New router
	InitRoute(e)
}

func InitRoute(e *echo.Echo) {
	gv2 := e.Group("arrow/api/v1.0/")

	gv2.GET("communicate", mAccountController.Communicate)
}
