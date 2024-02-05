package handler

import "booking-design/internal/core"

type handler struct {
	core *core.Core
}

func New(core *core.Core) *handler {
	return &handler{core}
}
