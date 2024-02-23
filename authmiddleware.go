package main

import (
	"fmt"

	"github.com/Saakhr/rssaggr/internal/auth"
	"github.com/Saakhr/rssaggr/internal/database"
	"github.com/labstack/echo/v4"
)

type authHandler func(echo.Context, database.User) error

func (apiCFG *apiconfig) middlewareAuth(handler authHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		apikey, err := auth.GetApiKey(c)
		if err != nil {
			return handleAnError(c, fmt.Sprintf("Error: %v", err))
		}

		user, err := apiCFG.DB.GetUserbyApiKey(c.Request().Context(), apikey)
		if err != nil {
			return handleAnError(c, fmt.Sprintf("couldn't get user: %v", err))
		}

		return handler(c, user)
	}

}
