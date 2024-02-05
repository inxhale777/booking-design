package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"booking-design/internal/api"
	"booking-design/internal/api/handler"
	"booking-design/internal/core"
	"booking-design/internal/entity"
	"booking-design/internal/entity/room"
	"booking-design/internal/sl"
	"booking-design/internal/usecase/booker"
	"booking-design/internal/usecase/bookingsaver"
	"booking-design/internal/usecase/provider"
)

func TestBookRoute(t *testing.T) {
	logger := sl.New()
	hryukingcom := provider.New(logger)
	saver := bookingsaver.New()

	c := core.Core{
		Providers: []entity.Provider{hryukingcom},
		Booker:    booker.New(saver),
		Logger:    logger,
	}

	router := api.Create(&c)

	now := time.Now()
	request := handler.BookRequest{
		Price:    "777",
		Provider: hryukingcom.Name(),
		Size:     string(room.Single),
		Category: string(room.Economy),
		From:     now.Add(24 * time.Hour),
		To:       now.Add(3 * 24 * time.Hour),
	}

	b, err := json.Marshal(request)
	require.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	var response handler.BookResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.Nil(t, err)

	require.Equal(t, 200, w.Code)
	require.NotEmpty(t, response.BookingID)
}
