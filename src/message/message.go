package message

import "fmt"

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
)

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
