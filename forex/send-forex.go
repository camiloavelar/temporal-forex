package forex

import (
	"context"

	"go.temporal.io/sdk/client"
)

func (f Forex) StartForex(ctx context.Context) error {
	options := client.StartWorkflowOptions{
		TaskQueue: "task-queue",
	}

	_, err := f.temporalClient.ExecuteWorkflow(ctx, options, "ForexWorkflow")

	return err
}

