package slot

import (
	"time"

	"booking-design/internal/entity/room"
)

type Request struct {
	Customer string
	Size     room.Size
	Category room.Category
	Price    string
	From     time.Time
	To       time.Time
}

type Slot struct {
	Count    int
	Size     room.Size
	Category room.Category
	Price    string
}
