package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type File struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FileName  string    `json:"file_name" db:"file_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
