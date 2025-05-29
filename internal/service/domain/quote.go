package domain

import (
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
)

type Quote struct {
	ID     int
	Author string
	Quote  string
}

func (q *Quote) ToModel() models.Quote {
	return models.Quote{
		Author: q.Author,
		Quote:  q.Quote,
	}
}
