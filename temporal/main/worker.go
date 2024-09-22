package main

import (
	"context"
	"log"
	"temporalavenue/account"
	"temporalavenue/forex"
	"temporalavenue/limit"
	messagebroker "temporalavenue/message-broker"
	"temporalavenue/temporal"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// Starting messageBroker and services
	messageBroker := messagebroker.NewBroker()
	limitActivities := limit.NewLimit(messageBroker, c)
	forexActivities := forex.NewForex(messageBroker, c)
	account := account.NewAccount(messageBroker, c)

	ctx := context.Background()

	// Listening and routing messages
	go func() {
		for {
			message := messageBroker.GetMessage()

			switch message.MessageType {
			case messagebroker.CheckLimitMessage:
				limitActivities.CheckLimit(ctx, message.Message, message.Token)
			case messagebroker.DebitAccountMessage:
				account.Debit(ctx, message.Message, message.Token)
			case messagebroker.DoForexMessage:
				forexActivities.DoForex(ctx, message.Message, message.Token)
			case messagebroker.CreditUSAccountMessage:
				account.Credit(ctx, message.Message, message.Token)
			}
		}
	}()

	w := worker.New(c, "task-queue", worker.Options{})
	w.RegisterWorkflow(temporal.ForexWorkflow)
	w.RegisterActivity(limitActivities)
	w.RegisterActivity(forexActivities)
	w.RegisterActivity(forex.IsValidAccount)
	w.RegisterActivity(forex.IsValidQuotation)
	w.RegisterActivity(forex.IsBusinessDay)
	w.RegisterActivity(forex.MarketOpen)
	w.RegisterActivity(forex.IsForexEnabled)
	w.RegisterActivity(account)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
