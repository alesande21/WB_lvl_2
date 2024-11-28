package entity

import "time"

type Event struct {
	ID     string    `json:"id"`
	UserID string    `json:"user_id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
}

type Events []Event
