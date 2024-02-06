package controller

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"akawork.io/dto" // Step 3
	"akawork.io/gerror"
	"akawork.io/infrastructure/controller" // Step 3
	"akawork.io/infrastructure/response"
	"akawork.io/infrastructure/util"

	// Step 3
	"github.com/labstack/echo"
)

/**
 * Defines a AccountController
 */
type HomePageController struct {
	controller.BaseController
}

/**
* Get an Account
 */
func (controller *HomePageController) Greeting(c echo.Context) error {
	vcolor := dto.ColorDto{}
	vcolor.Flag = "Hello member of training Docker, K8s organize by xPlat Team!!!!!!"
	return controller.WriteSuccess(c, vcolor)
}

func (controller *HomePageController) User(c echo.Context) error {

	out, err := exec.Command("whoami").Output()
	if err != nil {
		err := errors.New("Error run command inside container")
		gerr := gerror.New(gerror.ErrorSaveData, err, util.FuncName())
		message, resp := response.NewGerrorResponse(gerr)
		return controller.WriteInternalServerError(c, message, resp)
	}
	outx := strings.TrimSuffix(string(out), "\n")
	fmt.Println(outx)
	if outx == "root" {
		vcolor := dto.ColorDto{}
		vcolor.Flag = "Good luck to you next time!!!!!!"
		return controller.WriteSuccess(c, vcolor)
	}
	vcolor := dto.ColorDto{}
	vcolor.Flag = "Congratulation!!!!!!"
	return controller.WriteSuccess(c, vcolor)
}

func (controller *HomePageController) Communicate(c echo.Context) error {
	// Step 1
	// vcolor := dto.ColorDto{}
	// vcolor.Color = "Blue"
	// return controller.WriteSuccess(c, vcolor)

	// Step 2
	vcolor := dto.ColorDto{}
	vcolor.Flag = "Hello member of training Docker, K8s organize by xPlat Team!!!!!!"
	return controller.WriteSuccess(c, vcolor)

	// Step 3
	// err := errors.New("Error save data to Redis, need to use Key not List")
	// gerr := gerror.New(gerror.ErrorSaveData, err, util.FuncName())
	// message, resp := response.NewGerrorResponse(gerr)
	// return controller.WriteInternalServerError(c, message, resp)

}

func (controller *HomePageController) Parent(c echo.Context) error {
	// Step 1
	// vcolor := dto.ColorDto{}
	// vcolor.Color = "Blue"
	// return controller.WriteSuccess(c, vcolor)

	// Step 2
	vcolor := dto.ColorDto{}
	vcolor.Flag = "Hello member of training Docker, K8s organize by xPlat Team!!!!!!"
	return controller.WriteSuccess(c, vcolor)

	// Step 3
	// err := errors.New("Error save data to Redis, need to use Key not List")
	// gerr := gerror.New(gerror.ErrorSaveData, err, util.FuncName())
	// message, resp := response.NewGerrorResponse(gerr)
	// return controller.WriteInternalServerError(c, message, resp)

}

func (controller *HomePageController) GetHealthcheck(c echo.Context) error {

	vstatus := dto.HealthCheckDto{}
	vstatus.Status = "OK"
	// ecolor := response.ErrorResponse{}
	// ecolor.ErrorCode =
	// ecolor.Message =
	// ecolor.Exception =
	// err := errors.New("Error save data to Redis, need to use Key not List")
	// gerr := gerror.New(gerror.ErrorSaveData, err, util.FuncName())
	// message, resp := response.NewGerrorResponse(gerr)
	// return controller.WriteInternalServerError(c, message, resp)
	return controller.WriteSuccess(c, vstatus)
}
