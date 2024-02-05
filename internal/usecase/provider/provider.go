package provider

import (
	"context"
	"fmt"
	"sync"
	"time"

	"booking-design/internal/entity/slot"
	"booking-design/internal/randstr"
	"booking-design/internal/sl"
)

type freeroom struct {
	size     string
	category string
	from     time.Time
	to       time.Time
	count    int
}

var mu = sync.Mutex{}
var data = []freeroom{
	{
		size:     "single",
		category: "economy",
		from:     time.Now(),
		to:       time.Now().Add(24 * 7 * time.Hour),
		count:    7,
	},
	{
		size:     "triple",
		category: "lux",
		from:     time.Now().Add(24 * time.Hour),
		to:       time.Now().Add(24 * 3 * time.Hour),
		count:    1,
	},
}

type provider struct {
	logger *sl.Logger
}

func New(logger *sl.Logger) *provider {
	return &provider{
		logger,
	}
}

func (p *provider) Name() string {
	return "hryookingcom"
}

func (p *provider) Book(_ context.Context, req slot.Request) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	for i := range data {
		if data[i].category == string(req.Category) {
			if data[i].size == string(req.Size) {
				if req.From.After(data[i].from) && req.To.Before(data[i].to) {
					if data[i].count > 0 {
						// available room found
						data[i].count--

						p.logger.Info("available room found", "available", data[i].count)

						return randstr.Random(11), nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("no_room_available")
}

func (p *provider) Available(_ context.Context, _, _ time.Time) ([]slot.Slot, error) {
	// TODO implement me
	panic("implement me")
}
