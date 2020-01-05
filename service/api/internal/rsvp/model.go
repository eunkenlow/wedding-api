package rsvp

// Rsvp structure
type Rsvp struct {
	ID                  int    `json:"id" validate:"isdefault"`
	FirstName           string `json:"first_name" validate:"required"`
	LastName            string `json:"last_name" validate:"required"`
	Email               string `json:"email" validate:"required,email"`
	PhoneNumber         string `json:"phone_number" validate:"required"`
	Attending           *bool  `json:"attending"`
	Notifications       *bool  `json:"notifications"`
	DietaryRequirements string `json:"dietary_requirements"`
	BestMoments         string `json:"best_moments"`
}

// Rsvps represents list of rsvps
type Rsvps []*Rsvp
