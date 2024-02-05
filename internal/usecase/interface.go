package usecase

import (
	"context"

	"booking-design/internal/entity"
)

type (
	Booker interface {
		Book(ctx context.Context, req entity.BookRequest) (entity.Booking, error)
	}

	BookingSaver interface {
		Save(ctx context.Context, booking entity.Booking) (entity.Booking, error)
	}
)
