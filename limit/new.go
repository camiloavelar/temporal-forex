package limit

import messagebroker "temporalavenue/message-broker"

type Limit struct {
	messageBroker messagebroker.Broker
}

func NewLimit(messageBroker messagebroker.Broker) Limit {
	return Limit{
		messageBroker: messageBroker,
	}
}

