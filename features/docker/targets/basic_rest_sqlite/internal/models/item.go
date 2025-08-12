// Package models defines application data structures.
package models

import "time"

// Item is a simple record stored in the SQLite database.
type Item struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
