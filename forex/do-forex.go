package forex

import (
	"log/slog"
	messagebroker "temporalavenue/message-broker"
)

func (f Forex) DoForex(forex string) {
	slog.Info("received do forex message", "message", forex)
	slog.Info("sending forex to central bank...")
	slog.Info("✅ forex sent to central bank, now crediting US account")

	f.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.CreditUSAccountMessage,
		Message:     "please credit US account",
	})
}
