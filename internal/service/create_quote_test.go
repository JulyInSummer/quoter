package service

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
	mock_storage "github.com/JulyInSummer/quoter_app/internal/storage/mock"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func TestQuoter_CreateQuote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockRepoI(ctrl)
	req := domain.Quote{
		ID:     1,
		Author: "July",
		Quote:  "No Pain! No Gain!",
	}

	resp := &models.Quote{
		ID:     1,
		Author: "July",
		Quote:  "No Pain! No Gain!",
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	storage.EXPECT().CreateQuote(gomock.Any(), req.ToModel()).Return(resp, nil)

	service := NewQuoteService(logger, storage)

	res, err := service.CreateQuote(context.Background(), req)
	assert.NoError(t, err)

	assert.Equal(t, &req, res)
}
