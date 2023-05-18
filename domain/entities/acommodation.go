package entities

import "time"

type Acommodation struct {
	UserID   int       `json:"userId"`
	RoomID   int       `json:"roomId"`
	CheckIn  time.Time `json:"checkIn"`
	Checkout time.Time `json:"checkOut"`
	Garage   bool      `json:"garage"`
}

type Room struct {
	ID        int  `json:"id"`
	Available bool `json:"available"`
}
