package service

import (
	"context"
	"database/sql"
	"github.com/JulyInSummer/quoter_app/internal/config"
	mock_storage "github.com/JulyInSummer/quoter_app/internal/storage/mock"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func TestQuoter_Success_DeleteQuote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockRepoI(ctrl)

	quote := &models.Quote{
		ID:     1,
		Author: "July",
		Quote:  "No Pain! No Gain!",
	}

	storage.EXPECT().DeleteQuote(gomock.Any(), 1).Return(nil)
	storage.EXPECT().GetQuoteByID(gomock.Any(), 1).Return(quote, nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := NewQuoteService(logger, storage)

	err := service.DeleteQuote(context.Background(), 1)
	assert.NoError(t, err)
}

func TestQuoter_QuoteDoesNotExist_DeleteQuote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockRepoI(ctrl)

	storage.EXPECT().GetQuoteByID(gomock.Any(), 1).Return(nil, sql.ErrNoRows)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := NewQuoteService(logger, storage)

	err := service.DeleteQuote(context.Background(), 1)
	assert.Error(t, err)
	assert.ErrorIs(t, err, config.QuoteDoesNotExistError)
}
