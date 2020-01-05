package rsvp

type (
	// Service represents the service for rsvp.
	Service interface {
		CreateRsvp(rsvp *Rsvp) error
		FindAllRsvps() (Rsvps, error)
	}
	// ServiceImpl represents the service implementation for rsvp.
	ServiceImpl struct {
		repo Repository
	}
)

// NewService initialises a new rsvp service.
func NewService(repo Repository) *ServiceImpl {
	return &ServiceImpl{
		repo,
	}
}

// CreateRsvp creates new rsvp
func (s *ServiceImpl) CreateRsvp(rsvp *Rsvp) error {
	err := s.repo.CreateRsvp(rsvp)
	return err
}

// FindAllRsvps returns all rsvps
func (s *ServiceImpl) FindAllRsvps() (Rsvps, error) {
	rsvps, err := s.repo.FindAllRsvps()
	return rsvps, err
}
