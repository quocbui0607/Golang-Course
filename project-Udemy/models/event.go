package models

import (
	"time"

	"github.com/Wong-bui/Udemy-project/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	database := db.GetInstance()
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := database.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() (events []Event, err error) {
	database := db.GetInstance()

	query := "SELECT * FROM events"
	result, err := database.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event Event
		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	database := db.GetInstance()

	query := "SELECT * FROM events WHERE id = ?"
	result := database.QueryRow(query, id)

	var event Event
	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (event Event) Update() error {
	database := db.GetInstance()

	query := `UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?`

	stmt, err := database.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	database := db.GetInstance()

	query := "DELETE FROM events WHERE id = ?"

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	database := db.GetInstance()

	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := database.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	database := db.GetInstance()

	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := database.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err

}
