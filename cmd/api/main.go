package main

import (
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"booking-design/internal/api"
	"booking-design/internal/core"
	"booking-design/internal/entity"
	"booking-design/internal/sl"
	"booking-design/internal/usecase/booker"
	"booking-design/internal/usecase/bookingsaver"
	"booking-design/internal/usecase/provider"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := sl.New()
	hryukingcom := provider.New(logger)
	saver := bookingsaver.New()

	c := &core.Core{
		Providers: []entity.Provider{hryukingcom},
		Booker:    booker.New(saver),
		Logger:    logger,
	}

	r := api.Create(c)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("api server exited", sl.Err(err))
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		logger.Fatal("unable to shutdown graceful", sl.Err(err))
	}
}
