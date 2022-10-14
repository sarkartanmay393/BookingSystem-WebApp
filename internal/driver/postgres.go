package driver

import (
	"context"
	"log"
	"time"

	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
)

// InsertReservation inserts a new reservation into database.
func (db *DB) InsertReservation(res *models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO 
		reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	var newID int
	err := db.SQL.QueryRowContext(ctx, query, res.FirstName, res.LastName, res.Email, res.Phone, res.StartDate, res.EndDate, res.RoomID, time.Now(), time.Now()).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction inserts a new rooom restriction into database.
func (db *DB) InsertRoomRestriction(rr *models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO 
		room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.SQL.ExecContext(ctx, query, rr.StartDate, rr.EndDate, rr.RoomID, rr.ReservationID, rr.RestrictionID, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

// SearcAvailabilityByDatesAndRoomID return if room is availble in between selected days or not.
func (db *DB) SearcAvailabilityByDatesAndRoomID(start_date time.Time, end_date time.Time, room_id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT count(id) FROM room_restrictions WHERE 
	room_id = $1 and $2 < end_date and $3 > start_date;`

	var returnedCount int

	row := db.SQL.QueryRowContext(ctx, query, room_id, start_date, end_date)
	err := row.Scan(&returnedCount)
	if err != nil {
		return false, err
	}

	if returnedCount == 0 {
		return true, nil
	}
	return false, nil
}

// SearcAvailabilityByDates returns a slice of available rooms if any.
func (db *DB) SearcAvailabilityByDates(start_date, end_date time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, room_name FROM rooms 
	WHERE id NOT IN 
	(SELECT room_id 
	FROM room_restrictions 
	WHERE $1 < end_date and $2 > start_date);`

	roomsList := []models.Room{}

	rows, err := db.SQL.QueryContext(ctx, query, start_date, end_date)
	if err != nil {
		log.Println("Error while querying, Err:", err)
		return roomsList, err
	}
	for rows.Next() {
		var room models.Room
		rows.Scan(&room.ID, &room.RoomName)
		roomsList = append(roomsList, room)
	}

	if err = rows.Err(); err != nil {
		return roomsList, err
	}
	return roomsList, nil
}
