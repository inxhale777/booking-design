package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"booking-design/internal/api/handler"
	"booking-design/internal/core"
	"booking-design/internal/entity"
)

func Create(core *core.Core) http.Handler {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("user", entity.User{
			Email:     "johnwick@xyz.com",
			Phone:     "+3274648424",
			FirstName: "John",
			LastName:  "Wick",
		})
	})

	r.Use(func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		ctx.Next()

		status := ctx.Writer.Status()
		method := ctx.Request.Method
		host := ctx.Request.Host
		route := ctx.FullPath()
		end := time.Now()
		latency := end.Sub(start)
		ip := ctx.ClientIP()
		referer := ctx.Request.Referer()

		attrs := []slog.Attr{
			slog.Time("time", start),
			slog.String("method", method),
			slog.String("host", host),
			slog.String("path", path),
			slog.String("query", query),
			slog.String("route", route),
			slog.String("ip", ip),
			slog.String("referer", referer),
			slog.Time("time", end),
			slog.Duration("latency", latency),
			slog.Int("status", status),
		}

		core.Logger.LogAttrs(ctx.Request.Context(), slog.LevelInfo, "Incoming request", attrs...)
	})

	h := handler.New(core)
	r.POST("/book", h.Book)

	return r
}
