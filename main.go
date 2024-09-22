package main

import (
	"os"
	"os/signal"
	"syscall"
	"temporalavenue/account"
	"temporalavenue/forex"
	"temporalavenue/limit"
	messagebroker "temporalavenue/message-broker"
)

func main() {
	// Handling shutdown signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Starting messageBroker and services
	messageBroker := messagebroker.NewBroker()
	limit := limit.NewLimit(messageBroker)
	forex := forex.NewForex(messageBroker)
	account := account.NewAccount(messageBroker)

	// Listening and routing messages
	go func() {
		for {
			message := messageBroker.GetMessage()

			switch message.MessageType {
			case messagebroker.CheckLimitMessage:
				limit.CheckLimit(message.Message)
			case messagebroker.DebitAccountMessage:
				account.Debit(message.Message)
			case messagebroker.DoForexMessage:
				forex.DoForex(message.Message)
			case messagebroker.CreditUSAccountMessage:
				account.Credit(message.Message)
			}
		}
	}()

	// Start forex
	forex.StartForex()

	_ = <-sigs
}
