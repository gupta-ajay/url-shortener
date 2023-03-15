package db

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type JSONQueryHook struct {
	enabled bool
	verbose bool
	logger  *zap.Logger
}

var _ bun.QueryHook = (*JSONQueryHook)(nil)

func NewJSONQueryHook(logger *zap.Logger, enabled, verbose bool) *JSONQueryHook {
	h := &JSONQueryHook{
		enabled: enabled,
		verbose: verbose,
		logger:  logger,
	}
	return h
}

func (h *JSONQueryHook) BeforeQuery(
	ctx context.Context, event *bun.QueryEvent,
) context.Context {
	return ctx
}

func (h *JSONQueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	if !h.enabled {
		return
	}

	if !h.verbose {
		switch event.Err {
		case nil, sql.ErrNoRows, sql.ErrTxDone:
			return
		}
	}

	now := time.Now()
	dur := now.Sub(event.StartTime)
	m1 := regexp.MustCompile(`\\|\"`)

	logData := map[string]interface{}{
		"operation":      event.Operation(),
		"query":          m1.ReplaceAllString(event.Query, ""),
		"args":           event.QueryArgs,
		"model":          event.Model,
		"execution_time": fmt.Sprintf("%s", dur.Round(time.Microsecond)),
		"message":        "sqlQuery",
	}
	defer h.logger.Sync()

	if event.Err != nil {
		logData["error"] = event.Err.Error()
		logData["message"] = "sqlQueryErr"

		zapOptions := []zap.Option{zap.AddStacktrace(zap.PanicLevel)}
		h.logger.WithOptions(zapOptions...).Error("sqlQueryErr", zap.Any("data", logData))
		return
	}

	h.logger.Info("sqlQuery", zap.Any("data", logData))
}
