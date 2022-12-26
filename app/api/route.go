package api

import "github.com/labstack/echo"

func New(hdl *expenseHandler) *echo.Echo {
	e := echo.New()

	return e
}
