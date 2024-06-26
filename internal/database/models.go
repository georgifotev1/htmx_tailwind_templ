// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Category    string         `json:"category"`
	Location    string         `json:"location"`
	Quantity    int32          `json:"quantity"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Description sql.NullString `json:"description"`
	Price       sql.NullString `json:"price"`
	IsActive    sql.NullBool   `json:"is_active"`
}
