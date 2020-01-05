package score

import "github.com/go-pg/pg/v9"

type (
	// Repository represents the repository for score.
	Repository interface {
		CreateScore(score *Score) error
		FindAllScores() (Scores, error)
	}
	// RepositoryImpl represents the repository implementation for score.
	RepositoryImpl struct {
		db *pg.DB
	}
)

// NewRepository initialises a new score repository.
func NewRepository(db *pg.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// CreateScore creates a new score
func (r *RepositoryImpl) CreateScore(score *Score) error {
	err := r.db.Insert(score)
	return err
}

// FindAllScores returns all high scores
func (r *RepositoryImpl) FindAllScores() (Scores, error) {
	scores := Scores{}
	err := r.db.Model(&scores).
		Order("score DESC").
		Limit(10).
		Select()
	return scores, err
}
