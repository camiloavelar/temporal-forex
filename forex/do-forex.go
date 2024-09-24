package forex

import (
	"context"
	"log/slog"
	messagebroker "temporalavenue/message-broker"
	"time"

	"go.temporal.io/sdk/activity"
)

func (f *Forex) DoForex(ctx context.Context, forex string, taskToken []byte) error {
	slog.Info("received do forex message", "message", forex)
	slog.Info("sending forex to central bank...")
	slog.Info("✅ forex sent to central bank, now crediting US account")

	time.Sleep(2 * time.Second)

	return f.temporalClient.CompleteActivity(ctx, taskToken, nil, nil)
}

func (f *Forex) RequestForex(ctx context.Context, forex string) error {
	activityInfo := activity.GetInfo(ctx)
	taskToken := activityInfo.TaskToken

	f.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.DoForexMessage,
		Message:     "send forex to central bank",
		Token:       taskToken,
	})

	return activity.ErrResultPending
}
