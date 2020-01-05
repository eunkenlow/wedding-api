package score

type (
	// Service represents the service for score.
	Service interface {
		CreateScore(score *Score) error
		FindAllScores() (Scores, error)
	}
	// ServiceImpl represents the service implementation for score.
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

// CreateScore creates new score
func (s *ServiceImpl) CreateScore(score *Score) error {
	err := s.repo.CreateScore(score)
	return err
}

// FindAllScores returns all scores
func (s *ServiceImpl) FindAllScores() (Scores, error) {
	scores, err := s.repo.FindAllScores()
	return scores, err
}
