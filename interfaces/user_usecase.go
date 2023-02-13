package interfaces

import (
	"context"

	"gitlab.engdb.com.br/apigin/domain/entities"
)

type ScheduleUseCase interface {
	CreateSchedule(ctx context.Context, body entities.InputSchedule) (int, *entities.Error)

	GetSchedules(ctx context.Context) ([]entities.ResponseGetSchudeles, int, *entities.Error)

	CheckScheduleAvailability(ctx context.Context) ([]entities.ResponseAvailability, int, *entities.Error)
}
