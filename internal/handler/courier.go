package handler

import (
	"github.com/barancanatbas/courierTracking/domain/dto/request"
	"github.com/barancanatbas/courierTracking/internal/usecase"
	"github.com/barancanatbas/courierTracking/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CourierHandler interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
	Get(c echo.Context) error
}

type courier struct {
	util           *util.Util
	courierUsecase usecase.CourierUsecase
}

func NewCourierHandler(courierUsecase usecase.CourierUsecase, util *util.Util) CourierHandler {
	return &courier{
		util:           util,
		courierUsecase: courierUsecase,
	}
}

func (ch *courier) Create(c echo.Context) error {
	var req request.Courier
	if err := ch.util.Validate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := ch.courierUsecase.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (ch *courier) Delete(c echo.Context) error {
	return nil
}

func (ch *courier) Update(c echo.Context) error {
	return nil
}

func (ch *courier) Get(c echo.Context) error {
	return nil
}
