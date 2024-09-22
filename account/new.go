package account

import (
	messagebroker "temporalavenue/message-broker"

	"go.temporal.io/sdk/client"
)

type Account struct {
	messageBroker messagebroker.Broker
	temporalClient client.Client
}

func NewAccount(messageBroker messagebroker.Broker, client client.Client) *Account {
	return &Account{
		messageBroker: messageBroker,
		temporalClient: client,
	}
}

