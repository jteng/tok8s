package mock

import "cloud.google.com/go/pubsub"

type MessageItem struct {
	*pubsub.Message
}

func (m *MessageItem) Ack() {
}

func (m *MessageItem) Nack() {
}

func (m *MessageItem) Attributes() map[string]string {
	return m.Message.Attributes
}

func (m *MessageItem) Data() []byte {
	return m.Message.Data
}

func (m *MessageItem) ID() string {
	return m.Message.ID
}

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
