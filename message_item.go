package tok8s

import "cloud.google.com/go/pubsub"

//MessageItem implements Message interface, it is a wrapper around *pubsub.Message
type MessageItem struct {
	*pubsub.Message
}

//Ack acknowledges the message
func (m *MessageItem) Ack() {
	m.Message.Ack()
}

//Nack de-acknowledges the message
func (m *MessageItem) Nack() {
	m.Message.Nack()
}

//Attributes returns the attributes of the message
func (m *MessageItem) Attributes() map[string]string {
	return m.Message.Attributes
}

//Data returns the data content of the message
func (m *MessageItem) Data() []byte {
	return m.Message.Data
}

//ID returns the ID of the message
func (m *MessageItem) ID() string {
	return m.Message.ID
}

//SetData sets the data of the message
func (m *MessageItem) SetData(d []byte) {
	m.Message.Data = d
}

//NewMessageItem creates an instance of MessageItem with Attributes and Data initialized
func NewMessageItem() *MessageItem {
	return &MessageItem{
		&pubsub.Message{
			Attributes: make(map[string]string),
			Data:       make([]byte, 0),
		},
	}
}
