package forex

import (
	"context"
	"log/slog"
	"time"
)

func IsForexEnabled(_ context.Context) (bool, error) {
	slog.Info("forex is enabled")
	time.Sleep(2 * time.Second)
	return true, nil
}

func MarketOpen(_ context.Context) (bool, error) {
	slog.Info("market is open")
	time.Sleep(2 * time.Second)
	return true, nil
}

func IsBusinessDay(_ context.Context) (bool, error) {
	slog.Info("is business day")
	time.Sleep(2 * time.Second)
	return true, nil
}

func IsValidQuotation(_ context.Context) (bool, error) {
	slog.Info("quotation is valid")
	time.Sleep(2 * time.Second)
	return true, nil
}

func IsValidAccount(_ context.Context) (bool, error) {
	slog.Info("account is valid")
	time.Sleep(2 * time.Second)
	return true, nil
}
