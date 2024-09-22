package main

import (
	"context"
	"log"
	"log/slog"
	"temporalavenue/temporal"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: "task-queue",
	}

	_, err = c.ExecuteWorkflow(context.Background(), options, temporal.ForexWorkflow)
	if err != nil {
		slog.Error("error executing workflow", slog.Any("error", err))
	}
}
