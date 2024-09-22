package forex

import (
	"context"
	"log/slog"
)

func IsForexEnabled(_ context.Context) (bool, error) {
	slog.Info("forex is enabled")
	return true, nil
}

func MarketOpen(_ context.Context) (bool, error) {
	slog.Info("market is open")
	return true, nil
}

func IsBusinessDay(_ context.Context) (bool, error) {
	slog.Info("is business day")
	return true, nil
}

func IsValidQuotation(_ context.Context) (bool, error) {
	slog.Info("quotation is valid")
	return true, nil
}

func IsValidAccount(_ context.Context) (bool, error) {
	slog.Info("account is valid")
	return true, nil
}
