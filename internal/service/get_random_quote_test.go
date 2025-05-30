package service

import (
	"context"
	mock_storage "github.com/JulyInSummer/quoter_app/internal/storage/mock"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func TestQuoter_GetRandomQuote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockRepoI(ctrl)

	resp := &models.Quote{
		ID:     2,
		Author: "Walter White",
		Quote:  "I'm not in danger! I'M THE DANGER!!!",
	}

	storage.EXPECT().GetRandomQuote(gomock.Any()).Return(resp, nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := NewQuoteService(logger, storage)

	quote, err := service.GetRandomQuote(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, resp.ID, quote.ID)
	assert.Equal(t, resp.Author, quote.Author)
	assert.Equal(t, resp.Quote, quote.Quote)
}
