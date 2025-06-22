package models

import (
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
