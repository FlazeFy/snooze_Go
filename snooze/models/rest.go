package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Rest is used by pop to map your rests database table to your go code.
type Rest struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	RestNotes    *string    `json:"rest_notes" db:"rest_notes"`
	RestCategory string     `json:"rest_category" db:"rest_category"`
	StartedAt    time.Time  `json:"started_at" db:"started_at"`
	EndAt        *time.Time `json:"end_at" db:"end_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
}

func (r Rest) TableName() string {
	return "rests"
}

// String is not required by pop and may be deleted
func (r Rest) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Rests is not required by pop and may be deleted
type Rests []Rest

// String is not required by pop and may be deleted
func (r Rests) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Rest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Rest) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Rest) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
