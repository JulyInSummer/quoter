package resources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateQuoteRequest_ToDomain(t *testing.T) {
	req := CreateQuoteRequest{
		Author: "Walter White",
		Quote:  "I'm not in danger! I'M THE DANGER!!!",
	}

	domain := req.ToDomain()
	assert.Equal(t, req.Author, domain.Author)
	assert.Equal(t, req.Quote, domain.Quote)
}

func TestCreateQuoteRequest_Success_Validate(t *testing.T) {
	req := CreateQuoteRequest{
		Author: "Walter White",
		Quote:  "I'm not in danger! I'M THE DANGER!!!",
	}

	errs := req.Validate()
	assert.Len(t, errs, 0)
}

func TestCreateQuoteRequest_Required_Validate(t *testing.T) {
	req := CreateQuoteRequest{
		Quote: "I'm not in danger! I'M THE DANGER!!!",
	}

	errs := req.Validate()
	assert.Len(t, errs, 1)

	assert.Equal(t, "the field author is required", errs[0])
}
