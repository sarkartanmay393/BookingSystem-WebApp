package models

// Reservation struct holds data of a person who wants to reserve a room.
type Reservation struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

// ChosenDates structs for two dates to select together.
type ChosenDates struct {
	Start string
	End   string
}
