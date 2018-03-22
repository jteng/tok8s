package tok8s

import "strings"

//BuildMessage is the message sent by google container builder
type BuildMessage struct {
	Action string `json:"action"`
	Digest string `json:"digest"`
	Tag    string `json:"tag"`
}

//GetImage returns the image without SHA1
func (m BuildMessage) GetImage() string {
	if idx := strings.Index(m.Digest, ":"); idx != -1 {
		return m.Digest[:idx]
	}
	return m.Digest
}
