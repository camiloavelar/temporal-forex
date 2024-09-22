package forex

import (
	messagebroker "temporalavenue/message-broker"

	"go.temporal.io/sdk/client"
)

type Forex struct {
	messageBroker messagebroker.Broker
	temporalClient client.Client
}

func NewForex(messageBroker messagebroker.Broker, client client.Client) *Forex {
	return &Forex{
		messageBroker: messageBroker,
		temporalClient: client,
	}
}
