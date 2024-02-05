package entity

import (
	"context"
	"time"

	"booking-design/internal/entity/room"
	"booking-design/internal/entity/slot"
)

type User struct {
	Email     string
	Phone     string
	FirstName string
	LastName  string
}

type BookRequest struct {
	Price    string
	User     User
	Provider Provider
	Size     room.Size
	Category room.Category
	From     time.Time
	To       time.Time
}

type Booking struct {
	ID string

	Provider          string
	ProviderBookingID string

	Size     room.Size
	Category room.Category
	From     time.Time
	To       time.Time
}

type Provider interface {
	Name() string
	Book(ctx context.Context, req slot.Request) (string, error)
	Available(ctx context.Context, from, to time.Time) ([]slot.Slot, error)
}
