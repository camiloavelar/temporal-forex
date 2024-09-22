package forex

import "log/slog"

func IsForexEnabled() bool {
	slog.Info("forex is enabled")
	return true
}

func MarketOpen() bool {
	slog.Info("market is open")
	return true
}

func IsBusinessDay() bool {
	slog.Info("is business day")
	return true
}

func IsValidQuotation() bool {
	slog.Info("quotation is valid")
	return true
}

func IsValidAccount() bool {
	slog.Info("account is valid")
	return true
}
