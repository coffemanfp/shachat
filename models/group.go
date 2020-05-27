package models

import "time"

// Group - User group for the app.
type Group struct {
	ID int `json:"id,omitempty" db:"id,omitempty"`

	Link string `json:"link,omitempty" db:"link,omitempty"`
	Name string `json:"name,omitempty" db:"name,omitempty"`

	CreatedOn *time.Time `json:"createdOn,omitempty" db:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty" db:"updated_on,omitempty"`
}
