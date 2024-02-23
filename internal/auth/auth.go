package auth

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetApiKey(c echo.Context) (string, error) {
	val := c.Request().Header.Get("Authorization")

	if val == "" {
		return "", errors.New("no Authorization found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed api key")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part")
	}
	return vals[1], nil
}
