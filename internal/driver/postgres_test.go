package driver

import (
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"testing"
	"time"
)

func TestDB_InsertReservation(t *testing.T) {
	var res = &models.Reservation{
		FirstName: "Test",
		LastName:  "Person",
		Phone:     "09987-654-1234",
		Email:     "test@unit.in",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(24 * time.Hour),
		RoomID:    0,
	}

	_, err := db.InsertReservation(res)
	if err != nil {
		t.Error(err)
	}
}

func TestDB_InsertRoomRestriction(t *testing.T) {
	var res *models.RoomRestriction
	res = &models.RoomRestriction{
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		RoomID:        0,
		ReservationID: 0,
		RestrictionID: 0,
	}
	err := db.InsertRoomRestriction(res)
	if err != nil {
		t.Error(err)
	}
}

func TestDB_SearchAvailabilityByDatesByRoomID(t *testing.T) {
	_, err := db.SearchAvailabilityByDatesAndRoomID(time.Now(), time.Now(), 0)
	if err != nil {
		t.Error(err)
	}
}

func TestDB_SearchAvailabilityByDates(t *testing.T) {
	_, err := db.SearchAvailabilityByDates(time.Now(), time.Now())
	if err != nil {
		t.Error(err)
	}
}
