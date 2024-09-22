package limit

import (
	"context"
	"log/slog"
	messagebroker "temporalavenue/message-broker"

	"go.temporal.io/sdk/activity"
)

func (l *Limit) CheckLimit(ctx context.Context, limit string, taskToken []byte) error {
	slog.Info("received limit message", "message", limit)
	slog.Info("✅ limit is ok, starting to debit BR account")

	return l.temporalClient.CompleteActivity(ctx, taskToken, true, nil)
}

func (l *Limit) RequestCheckLimit(ctx context.Context) error {
	activityInfo := activity.GetInfo(ctx)
	taskToken := activityInfo.TaskToken

	slog.Info("sending message to check limit")
	l.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.CheckLimitMessage,
		Message:     "please check limit",
		Token:       taskToken,
	})

	return activity.ErrResultPending
}
