package usecase

import (
	"context"
	"github.com/barancanatbas/courierTracking/domain/dto/request"
	"github.com/barancanatbas/courierTracking/domain/dto/response"
	"github.com/barancanatbas/courierTracking/internal/repository"
	"github.com/barancanatbas/courierTracking/mapper"
	"github.com/barancanatbas/courierTracking/util"
)

type CourierUsecase interface {
	Create(ctx context.Context, req request.Courier) (response.Courier, error)
}

type courier struct {
	mapper     mapper.CourierMapper
	repository repository.Repository
	util       *util.Util
}

func NewCourierService(mapper mapper.CourierMapper, repository repository.Repository, util *util.Util) CourierUsecase {
	return &courier{
		mapper:     mapper,
		repository: repository,
		util:       util,
	}
}

func (cs *courier) Create(ctx context.Context, req request.Courier) (response.Courier, error) {
	courierModel := cs.mapper.CourierDtoToModel(req)

	err := cs.repository.Create(ctx, courierModel)
	if err != nil {
		return response.Courier{}, err
	}

	return cs.mapper.CourierModelToDto(courierModel), nil
}
