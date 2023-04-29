package util

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Util struct {
	validator *validator.Validate
}

func NewUtil() *Util {
	return &Util{
		validator.New(),
	}
}

func (u *Util) Validate(c echo.Context, i interface{}) error {
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := u.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
