package controller

import (
	"net/http"

	"akawork.io/infrastructure/logger"
	"akawork.io/infrastructure/response"
	"akawork.io/infrastructure/util"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

/**
 * Define a BaseController
 */
type BaseController struct {
}

/**
 * Returns a success response
 */
func (controller *BaseController) WriteSuccess(c echo.Context, v interface{}) error {
	res := response.Response{
		Message: "Success",
		Data:    v,
	}

	// Log response
	// logger.Info(util.ToJSON(res))

	// Return
	return c.JSON(http.StatusOK, res)
}

/**
 * Returns a success response without content
 */
func (controller *BaseController) WriteSuccessEmptyContent(c echo.Context) error {
	res := response.Response{
		Message: "Success",
		Data:    nil,
	}

	// Log response
	// logger.Info(util.ToJSON(res))

	// Return
	return c.JSON(http.StatusOK, res)
}

/**
 * Returns an error response
 */
func (controller *BaseController) writeError(c echo.Context, statusCode int, message string, err response.ErrorResponse) error {
	res := response.Response{
		Message: message,
		Data:    err,
	}
	// Log error
	logger.Error(util.ToJSON(res))
	// Return
	return c.JSON(statusCode, res)
}

/**
 * Returns an error as bad request (client-side error)
 */
func (controller *BaseController) WriteBadRequest(c echo.Context, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(c, http.StatusBadRequest, message, errorRes)
}

/**
 * Redirect an error as internal server error (server-side error)
 */
func (controller *BaseController) WriteInternalServerError(c echo.Context, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(c, http.StatusInternalServerError, message, errorRes)
}

/**
 * Validates model before do something
 */
func (controller *BaseController) IsValid(m interface{}) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
