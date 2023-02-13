package interfaces

import (
	"context"

	"gitlab.engdb.com.br/apigin/domain/entities"
)

type ScheduleRepo interface {
	CreateSchedule(ctx context.Context, body entities.InputSchedule) (int, error)
	GetSchedules(ctx context.Context) ([]entities.ResponseGetSchudeles, int, error)
	GetAllHours(ctx context.Context) ([]entities.ResponseAvailability, int, error)
	IsHourAvailabe(ctx context.Context, body entities.InputSchedule) (bool, int, error)
}
