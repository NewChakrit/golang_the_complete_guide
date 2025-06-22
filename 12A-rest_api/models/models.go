package models

import (
	"example.com/rest_api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"` // for not required field
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event

func (e Event) Save() error {
	query := `INSERT INTO events(name ,description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)` // todo ? เพื่อแทนค่า column ต่างๆ ใน table ตามลำดับ

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // todo Exec(args ...any) รับ args หลายตัวเพื่อรองรับค่า query
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // auto increment ID ให้
	if err != nil {
		return err
	}

	e.ID = id

	events = append(events, Event{})

	return nil
}

func GetAllEvents() []Event {
	return events
}
