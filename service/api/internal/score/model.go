package score

// Score structure
type Score struct {
	Name  string `json:"name" validate:"required"`
	Score int    `json:"score" validate:"required"`
}

// Scores represents list of scores
type Scores []*Score
