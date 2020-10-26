package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Foobar struct
type Foobar struct {
	ID            uuid.UUID `db:"id" json:"id"`
	FoobarContent string    `db:"foobar_content" json:"foobar_content"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
