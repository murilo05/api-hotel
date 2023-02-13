package usecase

import (
	"context"
	"errors"

	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
	errorUtils "gitlab.engdb.com.br/apigin/utils/error"
)

type scheduleUsecase struct {
	scheduleRepo interfaces.ScheduleRepo
}

func NewScheduleUsecase(us interfaces.ScheduleRepo) interfaces.ScheduleUseCase {
	return &scheduleUsecase{
		scheduleRepo: us,
	}
}

func (u *scheduleUsecase) GetSchedules(ctx context.Context) ([]entities.ResponseGetSchudeles, int, *entities.Error) {

	schedules, code, err := u.scheduleRepo.GetSchedules(ctx)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return nil, code, &errResp
	}

	return schedules, code, nil
}

func (u *scheduleUsecase) CreateSchedule(ctx context.Context, body entities.InputSchedule) (int, *entities.Error) {

	ok, code, err := u.scheduleRepo.IsHourAvailabe(ctx, body)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return code, &errResp
	}

	if !ok {
		errResp := errorUtils.CreateError(code, errors.New("horario indisponivel").Error())
		return code, &errResp
	}

	code, errors := u.scheduleRepo.CreateSchedule(ctx, body)
	if errors != nil {
		errResp := errorUtils.CreateError(code, errors.Error())
		return code, &errResp
	}

	return 200, nil
}

func (u *scheduleUsecase) CheckScheduleAvailability(ctx context.Context) ([]entities.ResponseAvailability, int, *entities.Error) {

	ScheduleAvailability, code, err := u.scheduleRepo.GetAllHours(ctx)
	if err != nil {
		errResp := errorUtils.CreateError(code, err.Error())
		return nil, code, &errResp
	}

	outTransformed := []entities.ResponseAvailability{}
	var available bool

	for _, x := range ScheduleAvailability {

		if x.AvailableFromSQL == 1 {
			available = true
		} else {
			available = false
		}

		outTransformed = append(outTransformed, entities.ResponseAvailability{
			StartingHour: x.StartingHour,
			FinalHour:    x.FinalHour,
			Available:    available,
		})
	}

	return outTransformed, 200, nil

}
