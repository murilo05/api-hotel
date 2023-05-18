package interfaces

import (
	"context"

	"api/api-hotel/domain/entities"
)

type HotelRepo interface {
	ListUsers(ctx context.Context, userInfo entities.User) (user []entities.User, err error)
	RegisterUser(ctx context.Context, user entities.User) (err error)
	UpdateUser(ctx context.Context, user entities.User, userID int) (err error)
	DeleteUser(ctx context.Context, userID int) (err error)
	RegisterReservation(ctx context.Context, acommodation entities.Acommodation, price float64) (err error)
	IsRoomEmpty(ctx context.Context, roomID int) (isRoomEmpty bool, err error)
	ListAcommodations(ctx context.Context) (acommodtions []entities.Acommodation, err error)
}
