package account

import (
	"context"
	"log/slog"
	messagebroker "temporalavenue/message-broker"

	"go.temporal.io/sdk/activity"
)

func (a *Account) Debit(ctx context.Context, debit string, taskToken []byte) error {
	slog.Info("receved debit call", "debit", debit)
	slog.Info("debiting account")
	slog.Info("✅ account debit ok, now sending forex to central bank")

	return a.temporalClient.CompleteActivity(ctx, taskToken, nil, nil)
}

func (a *Account) RequestDebit(ctx context.Context, debit string) error {
	activityInfo := activity.GetInfo(ctx)
	taskToken := activityInfo.TaskToken

	a.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.DebitAccountMessage,
		Message:     "please debit this account",
		Token:       taskToken,
	})

	return activity.ErrResultPending
}

func (a *Account) Credit(ctx context.Context, credit string, taskToken []byte) error {
	slog.Info("receved credit call", "credit", credit)
	slog.Info("crediting account")
	slog.Info("✅ account credit ok, forex is now completed")

	return a.temporalClient.CompleteActivity(ctx, taskToken, nil, nil)
}

func (a *Account) RequestCredit(ctx context.Context, debit string) error {
	activityInfo := activity.GetInfo(ctx)
	taskToken := activityInfo.TaskToken

	a.messageBroker.SendMessage(messagebroker.Message{
		MessageType: messagebroker.CreditUSAccountMessage,
		Message:     "please credit this account",
		Token:       taskToken,
	})

	return activity.ErrResultPending
}
