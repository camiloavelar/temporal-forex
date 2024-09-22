package limit

import (
	"log/slog"
	messagebroker "temporalavenue/message-broker"
)

func (l Limit) CheckLimit(limit string) {
	slog.Info("received limit message", "message", limit)
	slog.Info("✅ limit is ok, starting to debit BR account")

	l.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.DebitAccountMessage,
		Message:     "please debit this account",
	})
}
