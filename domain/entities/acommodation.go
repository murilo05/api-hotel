package entities

import "time"

type Acommodation struct {
	ID       int       `json:"Id"`
	UserID   int       `json:"userId"`
	RoomID   int       `json:"roomId"`
	CheckIn  time.Time `json:"checkIn"`
	Checkout time.Time `json:"checkOut"`
	Garage   bool      `json:"garage"`
	Price    float32   `json:"price"`
}

type Room struct {
	ID        int  `json:"id"`
	Available bool `json:"available"`
}
