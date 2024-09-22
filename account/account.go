package account

import (
	"log/slog"
	messagebroker "temporalavenue/message-broker"
)

func (a Account) Debit(debit string) {
	slog.Info("receved debit call", "debit", debit)
	slog.Info("debiting account")
	slog.Info("✅ account debit ok, now sending forex to central bank")

	a.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.DoForexMessage,
		Message: "please do Forex on central bank",
	})
}

func (a Account) Credit(credit string) {
	slog.Info("receved credit call", "credit", credit)
	slog.Info("crediting account")
	slog.Info("✅ account credit ok, forex is now completed")
}
