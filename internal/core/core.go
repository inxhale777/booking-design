package core

import (
	"booking-design/internal/entity"
	"booking-design/internal/sl"
	"booking-design/internal/usecase"
)

type Core struct {
	Providers []entity.Provider
	Booker    usecase.Booker
	Logger    *sl.Logger
}
