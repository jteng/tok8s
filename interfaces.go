package tok8s

//Subscriber is the interface
type Subscriber interface {
	Subscribe() <-chan Message
}

//Message interface
type Message interface {
	Ack()
	Nack()
	Attributes() map[string]string
	Data() []byte
	ID() string
}
