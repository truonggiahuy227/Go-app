package controller

import (
	"encoding/json"

	"akawork.io/dto" // Step 3
	"akawork.io/infrastructure/cache"
	"akawork.io/infrastructure/controller" // Step 3
	"akawork.io/module/account/service"

	// Step 3
	"github.com/labstack/echo"
)

/**
 * Defines a AccountController
 */
type AccountController struct {
	controller.BaseController
	Cache        cache.CacheManager
	CacheService service.ICacheService
}

/**
* Get an Account
 */
func (controller *AccountController) Communicate(c echo.Context) error {
	vcolor := dto.ColorDto{}
	controller.CacheService.CreatePostinCache()
	data := controller.CacheService.GetPostinCache()
	json.Unmarshal([]byte(data), &vcolor)
	return controller.WriteSuccess(c, vcolor)
}
