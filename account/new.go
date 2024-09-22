package account

import messagebroker "temporalavenue/message-broker"

type Account struct {
	messageBroker messagebroker.Broker
}

func NewAccount(messageBroker messagebroker.Broker) Account {
	return Account{
		messageBroker: messageBroker,
	}
}

