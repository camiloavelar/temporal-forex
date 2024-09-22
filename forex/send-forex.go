package forex

import (
	"log/slog"
	messagebroker "temporalavenue/message-broker"
)

func (f Forex) StartForex() {
	slog.Info("💸 starting forex validations")
	if !IsForexEnabled() {
		slog.Warn("forex is not enabled")
		return
	}

	if !MarketOpen() {
		slog.Warn("forex is not enabled")
		return
	}

	if !IsBusinessDay() {
		slog.Warn("is not business day")
		return
	}

	if !IsValidQuotation() {
		slog.Warn("quotation is not valid")
		return
	}

	if !IsValidAccount() {
		slog.Warn("account is not valid")
		return
	}

	slog.Info("✅ all validations are ok")

	f.CheckAccountLimit()
}

func (f Forex) CheckAccountLimit() {
	slog.Info("sending message to check limit")
	f.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.CheckLimitMessage,
		Message:     "please check limit",
	})
}
