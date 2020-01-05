package rsvp

import "github.com/go-pg/pg/v9"

type (
	// Repository represents the repository for rsvp.
	Repository interface {
		CreateRsvp(rsvp *Rsvp) error
		FindAllRsvps() (Rsvps, error)
	}
	// RepositoryImpl represents the repository implementation for rsvp.
	RepositoryImpl struct {
		db *pg.DB
	}
)

// NewRepository initialises a new rsvp repository.
func NewRepository(db *pg.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// CreateRsvp creates a new rsvp
func (r *RepositoryImpl) CreateRsvp(rsvp *Rsvp) error {
	err := r.db.Insert(rsvp)
	return err
}

// FindAllRsvps returns all rsvps
func (r *RepositoryImpl) FindAllRsvps() (Rsvps, error) {
	rsvps := Rsvps{}
	err := r.db.Model(&rsvps).Select()
	return rsvps, err
}
