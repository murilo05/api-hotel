package interfaces

import (
	"context"

	"api/api-hotel/domain/entities"
)

type ScheduleUseCase interface {
	ListUsers(ctx context.Context, userInfo entities.User) (user []entities.User, err error)

	RegisterUser(ctx context.Context, user entities.User) (err error)

	UpdateUser(ctx context.Context, user entities.User, userID int) (err error)

	DeleteUser(ctx context.Context, userID int) (err error)

	RegisterReservation(ctx context.Context, acommodation entities.Acommodation) (err error)
}
