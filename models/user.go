package models

import (
	"time"
)

// User - User of the app.
type User struct {
	ID int `json:"id,omitempty" db:"id,omitempty"`

	Link     string `json:"link,omitempty" db:"link,omitempty"`
	Password string `json:"password,omitempty" db:"password,omitempty"`

	CreatedOn *time.Time `json:"createdOn,omitempty" db:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty" db:"updated_on,omitempty"`
}
