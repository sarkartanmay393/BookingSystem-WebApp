package models

import "time"

// Reservation hold information of reservation performed by a user.
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

// ChosenDates structs for two dates to select together.
type ChosenDates struct {
	Start time.Time
	End   time.Time
}
