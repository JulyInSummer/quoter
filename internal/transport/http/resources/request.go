package resources

import (
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
)

type CreateQuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (r *CreateQuoteRequest) ToDomain() domain.Quote {
	return domain.Quote{
		Author: r.Author,
		Quote:  r.Quote,
	}
}

func (r *CreateQuoteRequest) Validate() []string {
	validationErrors := make([]string, 0)

	if r.Quote == "" {
		validationErrors = append(validationErrors, "the field quote is required")
	}

	if r.Author == "" {
		validationErrors = append(validationErrors, "the field author is required")
	}

	return validationErrors
}
