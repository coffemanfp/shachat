package models

// Contact - User contact.
type Contact struct {
	ID int `json:"id,omitempty" db:"db,omitempty"`

	User
}
