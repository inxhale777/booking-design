package booker

import (
	"context"
	"fmt"

	"booking-design/internal/entity"
	"booking-design/internal/entity/slot"
	"booking-design/internal/usecase"
)

type booker struct {
	saver usecase.BookingSaver
}

func New(saver usecase.BookingSaver) *booker {
	return &booker{
		saver,
	}
}

func (b *booker) Book(ctx context.Context, req entity.BookRequest) (entity.Booking, error) {
	const trace = "booker.Book"

	providerBookingID, err := req.Provider.Book(ctx, slot.Request{
		Customer: fmt.Sprintf("%s %s", req.User.FirstName, req.User.LastName),
		Size:     req.Size,
		Category: req.Category,
		Price:    req.Price,
		From:     req.From,
		To:       req.To,
	})
	if err != nil {
		return entity.Booking{}, fmt.Errorf("%s: %w", trace, err)
	}

	return b.saver.Save(ctx, entity.Booking{
		Provider:          req.Provider.Name(),
		ProviderBookingID: providerBookingID,
		Size:              req.Size,
		Category:          req.Category,
		From:              req.From,
		To:                req.To,
	})
}
