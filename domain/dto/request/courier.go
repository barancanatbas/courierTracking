package request

import "github.com/barancanatbas/courierTracking/domain/dto"

type Courier struct {
	Name        string          `json:"name" validate:"required"`
	Phone       string          `json:"phone" validate:"required"`
	VehicleType dto.VehicleType `json:"vehicleType" validate:"required"`
}
