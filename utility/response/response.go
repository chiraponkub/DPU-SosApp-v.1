package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func EchoSucceed(c echo.Context, msg interface{}) error {
	return c.JSON(http.StatusOK, msg)
}

func EchoError(c echo.Context, statusCode int, msg interface{}) error {
	return c.JSON(statusCode, msg)
}
