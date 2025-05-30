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

func TestQuoter_GetAllQuotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockRepoI(ctrl)

	quotes := []models.Quote{
		{
			ID:     1,
			Author: "July",
			Quote:  "No Pain! No Gain!",
		},
		{
			ID:     2,
			Author: "Walter White",
			Quote:  "I'm not in danger! I'M THE DANGER!!!",
		},
	}

	storage.EXPECT().GetAllQuotes(gomock.Any(), "").Return(quotes, nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := NewQuoteService(logger, storage)

	resp, err := service.GetAllQuotes(context.Background(), "")
	assert.NoError(t, err)
	assert.Len(t, resp, 2)

	assert.Equal(t, resp[0].ID, 1)
	assert.Equal(t, resp[0].Author, "July")
	assert.Equal(t, resp[0].Quote, "No Pain! No Gain!")

	assert.Equal(t, resp[1].ID, 2)
	assert.Equal(t, resp[1].Author, "Walter White")
	assert.Equal(t, resp[1].Quote, "I'm not in danger! I'M THE DANGER!!!")
}

func TestQuoter_ByAuthor_GetAllQuotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockRepoI(ctrl)

	quotes := []models.Quote{
		{
			ID:     2,
			Author: "Walter White",
			Quote:  "I'm not in danger! I'M THE DANGER!!!",
		},
	}

	storage.EXPECT().GetAllQuotes(gomock.Any(), "Walter White").Return(quotes, nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := NewQuoteService(logger, storage)

	resp, err := service.GetAllQuotes(context.Background(), "Walter White")
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	assert.Equal(t, resp[0].ID, 2)
	assert.Equal(t, resp[0].Author, "Walter White")
	assert.Equal(t, resp[0].Quote, "I'm not in danger! I'M THE DANGER!!!")
}
