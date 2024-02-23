package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleError(e echo.Context) error {
	er := errResponsed{
		Erorr: "Something Went Wrong!",
	}
	return e.JSON(http.StatusInternalServerError, er)
}
func handleHealth(e echo.Context) error {
	return e.JSON(http.StatusOK, struct{}{})
}
