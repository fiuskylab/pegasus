package topic

import (
	"github.com/fiuskylab/pegasus/src/message"
	"github.com/google/uuid"
)

type (
	// Topic is a "named queue" that services can subscribe to.
	Topic struct {
		input  chan *message.Message
		output chan *message.Message
		Name   string
	}
)

// NewTopic creates a new Topic, if given name is empty, it generates
// an UUID as a name.
func NewTopic(name string) *Topic {
	if name == "" {
		name = uuid.NewString()
	}

	return &Topic{
		input:  make(chan *message.Message),
		output: make(chan *message.Message),
		Name:   name,
	}
}

// Send will put a message in Topic input channel.
func (t *Topic) Send(m *message.Message) error {
	t.input <- m
	return nil
}

// Pop will retrieve and delete a message from the output channel.
func (t *Topic) Pop() (*message.Message, error) {
	return <-t.output, nil
}
