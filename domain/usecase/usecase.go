package usecase

import (
	"context"
	"errors"
	"time"

	"api/api-hotel/domain/entities"
	"api/api-hotel/interfaces"
)

type hotelUsecase struct {
	hotelRepo interfaces.HotelRepo
}

func NewHotelUsecase(us interfaces.HotelRepo) interfaces.ScheduleUseCase {
	return &hotelUsecase{
		hotelRepo: us,
	}
}

const (
	weekPrice          = 120.0
	weekEndPrice       = 150.0
	garageWeekPrice    = 15.0
	garageWeekEndPrice = 20.0
	extraDayPrice      = 1
	checkOutHour       = 16
	checkOutMinutes    = 30
	numRooms           = 50
)

func (u *hotelUsecase) ListUsers(ctx context.Context, userInfo entities.User) (users []entities.User, err error) {

	users, err = u.hotelRepo.ListUsers(ctx, userInfo)
	if err != nil {
		return
	}

	return
}

func (u *hotelUsecase) UpdateUser(ctx context.Context, user entities.User, userID int) (err error) {
	err = u.hotelRepo.UpdateUser(ctx, user, userID)
	return
}

func (u *hotelUsecase) DeleteUser(ctx context.Context, userID int) (err error) {
	err = u.hotelRepo.DeleteUser(ctx, userID)
	return
}

func (u *hotelUsecase) RegisterUser(ctx context.Context, user entities.User) (err error) {

	err = u.hotelRepo.RegisterUser(ctx, user)
	if err != nil {
		return
	}

	return nil

}

func (u *hotelUsecase) RegisterReservation(ctx context.Context, acommodation entities.Acommodation) (err error) {
	isRoomEmpty, err := u.hotelRepo.IsRoomEmpty(ctx, acommodation.RoomID)
	if err != nil {
		return
	}

	if !isRoomEmpty {
		err = errors.New("this room is already in use")
		return
	}

	price := u.calculatePrice(acommodation)

	err = u.hotelRepo.RegisterReservation(ctx, acommodation, price)

	return
}

func (u *hotelUsecase) calculatePrice(acommodation entities.Acommodation) float64 {
	numWeekdays := 0
	numWeekends := 0
	for d := acommodation.CheckIn; d.Before(acommodation.Checkout); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			numWeekends++
		} else {
			numWeekdays++
		}
	}

	price := float64(numWeekdays)*weekPrice + float64(numWeekends)*weekEndPrice
	if acommodation.Garage {
		price += float64(numWeekdays)*garageWeekPrice + float64(numWeekends)*garageWeekEndPrice
	}

	if acommodation.Checkout.Hour() >= checkOutHour && acommodation.Checkout.Minute() >= checkOutMinutes {
		price += weekPrice + weekEndPrice
	}

	return price
}
