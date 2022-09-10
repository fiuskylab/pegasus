package topic

import (
	"github.com/fiuskylab/pegasus/pkg/message"
	"github.com/google/uuid"
)

type (
	// Topic is a "named queue" that services can subscribe to.
	Topic struct {
		input  chan message.Message
		output chan message.Message
		Name   string
	}
)

// NewTopic creates a new Topic, if given name is empty, it generates
// an UUID as a name.
func NewTopic(name string) Topic {
	if name == "" {
		name = uuid.NewString()
	}

	return Topic{
		input:  make(chan message.Message),
		output: make(chan message.Message),
		Name:   name,
	}
}
