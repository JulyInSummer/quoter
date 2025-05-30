package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuote_ToModel(t *testing.T) {
	domain := Quote{
		Author: "July",
		Quote:  "No Pain! No Gain!",
	}

	res := domain.ToModel()

	assert.Equal(t, domain.ID, res.ID)
	assert.Equal(t, domain.Author, res.Author)
	assert.Equal(t, domain.Quote, res.Quote)
}
