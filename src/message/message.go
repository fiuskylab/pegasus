package message

import (
	"fmt"

	"github.com/fiuskylab/pegasus/src/proto"
)

type (
	// Message - A data model that represents the messages
	// sent/received to/from services subscribed to a Topic.
	Message struct {
		Body      string `json:"body"`
		TopicName string `json:"topic_name"`
		// TODO:
		// 	- Add more features, like:
		// 		- Attributes
		// 		- Expiration
		// 		- ID ???
	}
)

const (
	errEmptyField = "field '%s' is empty"
	errNilRequest = "received request is nil"
)

// FromRequest builds a Message from SendRequest and validates it.
func FromRequest(req *proto.SendRequest) (*Message, error) {
	if req == nil {
		return nil, fmt.Errorf(errNilRequest)
	}
	msg := &Message{
		Body:      req.Body,
		TopicName: req.TopicName,
	}

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	return msg, nil
}

// Validate - validates Message fields
func (m *Message) Validate() error {
	if m.Body == "" {
		return fmt.Errorf(errEmptyField, "body")
	}
	if m.TopicName == "" {
		return fmt.Errorf(errEmptyField, "topic_name")
	}
	return nil
}
