package mock

import (
	"encoding/json"

	"github.com/jteng/tok8s"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	Messages []tok8s.BuildMessage
}

func (s Subscriber) Subscribe() <-chan tok8s.Message {
	out := make(chan tok8s.Message)
	go func() {
		defer close(out)
		for _, m := range s.Messages {
			data, err := json.Marshal(m)
			if err != nil {
				logrus.Errorf("failed to marshal json %s", err.Error())
				continue
			}
			msg := NewMessageItem()
			msg.SetData(data)
			out <- msg
		}
	}()
	return out
}
