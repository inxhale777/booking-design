package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"booking-design/internal/entity"
	"booking-design/internal/entity/room"
)

type BookRequest struct {
	Price    string    `json:"price"`
	Provider string    `json:"provider"`
	Size     string    `json:"size"`
	Category string    `json:"category"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type BookResponse struct {
	BookingID string `json:"booking_id"`
}

func (h *handler) Book(ctx *gin.Context) {
	var request BookRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("malformed: %w", err))
		return
	}

	var provider entity.Provider
	for _, p := range h.core.Providers {
		if p.Name() == request.Provider {
			provider = p
			break
		}
	}

	if provider == nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid provider: %s", request.Provider))
		return
	}

	var size room.Size
	for _, s := range room.AllSizes {
		if string(s) == request.Size {
			size = s
			break
		}
	}

	var category room.Category
	for _, c := range room.AllCategories {
		if string(c) == request.Category {
			category = c
			break
		}
	}

	result, err := h.core.Booker.Book(ctx, entity.BookRequest{
		Price:    request.Price,
		User:     ctx.MustGet("user").(entity.User),
		Provider: provider,
		Size:     size,
		Category: category,
		From:     request.From,
		To:       request.To,
	})
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, BookResponse{
		BookingID: result.ID,
	})
}
