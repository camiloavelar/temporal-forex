package forex

import messagebroker "temporalavenue/message-broker"

type Forex struct {
	messageBroker messagebroker.Broker
}

func NewForex(messageBroker messagebroker.Broker) Forex {
	return Forex{
		messageBroker: messageBroker,
	}
}
