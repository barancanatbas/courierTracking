package mapper

import (
	"github.com/barancanatbas/courierTracking/domain/dto"
	"github.com/barancanatbas/courierTracking/domain/dto/request"
	"github.com/barancanatbas/courierTracking/domain/dto/response"
	"github.com/barancanatbas/courierTracking/domain/model"
)

type CourierMapper struct {
}

func NewCourierMapper() *CourierMapper {
	return &CourierMapper{}
}

func (cm *CourierMapper) CourierDtoToModel(req request.Courier) model.Courier {
	var result model.Courier

	result.Name = req.Name
	result.Phone = req.Phone
	result.VehicleType = model.VehicleType(req.VehicleType)

	return result
}

func (cm *CourierMapper) CourierModelToDto(req model.Courier) response.Courier {
	var result response.Courier

	result.ID = req.ID
	result.CreatedAt = req.CreatedAt
	result.Name = req.Name
	result.Phone = req.Phone
	result.VehicleType = dto.VehicleType(req.VehicleType)

	return result
}
