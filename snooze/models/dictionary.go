package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

// Dictionary is used by pop to map your dictionaries database table to your go code.
type Dictionary struct {
	ID      int    `json:"id" db:"id"`
	DctType string `json:"dictionary_type" db:"dictionary_type"`
	DctName string `json:"dictionary_name" db:"dictionary_name"`
}

func (r Dictionary) TableName() string {
	return "dictionaries"
}

// String is not required by pop and may be deleted
func (d Dictionary) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Dictionaries is not required by pop and may be deleted
type Dictionaries []Dictionary

// String is not required by pop and may be deleted
func (d Dictionaries) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Dictionary) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Dictionary) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Dictionary) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
