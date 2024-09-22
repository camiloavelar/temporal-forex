package limit

import (
	messagebroker "temporalavenue/message-broker"

	"go.temporal.io/sdk/client"
)

type Limit struct {
	messageBroker messagebroker.Broker
	temporalClient client.Client
}

func NewLimit(messageBroker messagebroker.Broker, client client.Client) *Limit {
	return &Limit{
		messageBroker: messageBroker,
		temporalClient: client,
	}
}

