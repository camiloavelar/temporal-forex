package messagebroker

type (
	Message struct {
		MessageType int
		Message     string
	}

	broker struct {
		channel chan Message
	}

	Broker interface {
		SendMessage(message Message) error
		GetMessage() Message
	}
)

const (
	CheckLimitMessage = iota
	DebitAccountMessage
	DoForexMessage
	CreditUSAccountMessage
)

func NewBroker() Broker {
	return broker{
		channel: make(chan Message, 1),
	}
}

func (b broker) SendMessage(message Message) error {
	b.channel <- message
	return nil
}

func (b broker) GetMessage() Message {
	return <-b.channel
}
