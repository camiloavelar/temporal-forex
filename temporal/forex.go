package temporal

import (
	"errors"
	"log/slog"
	"temporalavenue/account"
	"temporalavenue/forex"
	"temporalavenue/limit"
	"time"

	"go.temporal.io/sdk/workflow"
)

func ForexWorkflow(
	ctx workflow.Context,
) error {
	var (
		fxActivities      *forex.Forex
		accountActivities *account.Account
		limitActivities   *limit.Limit

		forexEnabled    bool
		marketOpen      bool
		businessDay     bool
		validQuotation  bool
		validAccount    bool
		accountHasLimit bool
	)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	slog.Info("💸 starting forex validations")
	_ = workflow.ExecuteActivity(ctx, forex.IsForexEnabled).Get(ctx, &forexEnabled)
	if !forexEnabled {
		return errors.New("forex is not enabled")
	}

	_ = workflow.ExecuteActivity(ctx, forex.MarketOpen).Get(ctx, &marketOpen)
	if !marketOpen {
		return errors.New("market is not open")
	}

	_ = workflow.ExecuteActivity(ctx, forex.IsBusinessDay).Get(ctx, &businessDay)
	if !businessDay {
		return errors.New("is not business day")
	}

	_ = workflow.ExecuteActivity(ctx, forex.IsValidQuotation).Get(ctx, &validQuotation)
	if !validQuotation {
		return errors.New("quotation is not valid")
	}

	_ = workflow.ExecuteActivity(ctx, forex.IsValidAccount).Get(ctx, &validAccount)
	if !validAccount {
		return errors.New("account is not valid")
	}

	slog.Info("✅ all validations are ok")

	_ = workflow.ExecuteActivity(ctx, limitActivities.RequestCheckLimit).Get(ctx, &accountHasLimit)
	if !accountHasLimit {
		return errors.New("account has no limit")
	}

	if err := workflow.ExecuteActivity(ctx, accountActivities.RequestDebit).Get(ctx, nil); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, fxActivities.RequestForex).Get(ctx, nil); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, accountActivities.RequestCredit).Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
