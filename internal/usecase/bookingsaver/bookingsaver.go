package bookingsaver

import (
	"context"

	"booking-design/internal/entity"
	"booking-design/internal/randstr"
)

type bookingsaver struct{}

func New() *bookingsaver {
	return &bookingsaver{}
}

func (bs *bookingsaver) Save(_ context.Context, booking entity.Booking) (entity.Booking, error) {
	return entity.Booking{
		ID:                randstr.Random(11),
		Provider:          booking.Provider,
		ProviderBookingID: booking.ProviderBookingID,
		Size:              booking.Size,
		Category:          booking.Category,
		From:              booking.From,
		To:                booking.To,
	}, nil
}
