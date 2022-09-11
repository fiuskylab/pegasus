package manager

import (
	"fmt"

	"github.com/fiuskylab/pegasus/src/message"
	"github.com/fiuskylab/pegasus/src/topic"
	"go.uber.org/zap"
)

type (
	// Manager - Manages Pegasus topics.
	Manager struct {
		topics map[string]*topic.Topic
	}
)

// NewManager returns a new Manager
func NewManager() *Manager {
	zap.L().Debug("creating a manager")
	return &Manager{
		topics: map[string]*topic.Topic{},
	}
}

// GetTopicNames returns a list of names of all created Topics.
func (m *Manager) GetTopicNames() []string {
	topics := []string{}

	for k := range m.topics {
		topics = append(topics, k)
	}

	return topics
}

const (
	errTopicAlreadyExists = "topic named '%s' already exists"
	errTopicNotFound      = "topic named '%s' not found"
)

// NewTopic creates a new Topic and adds it to the Manager.
func (m *Manager) NewTopic(name string) error {
	zap.L().Sugar().Infof("creating '%s' topic", name)
	if _, ok := m.topics[name]; ok {
		return fmt.Errorf(errTopicAlreadyExists, name)
	}
	createdTopic, err := topic.NewTopic(name)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	m.topics[createdTopic.Name] = createdTopic
	return nil
}

// Send inserts a message into Topic's internal queue.
func (m *Manager) Send(msg *message.Message) error {
	zap.L().Info("sending message", zap.Any("message", *msg))
	if err := msg.Validate(); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	topic, ok := m.topics[msg.TopicName]
	if !ok {
		return fmt.Errorf(errTopicNotFound, msg.TopicName)
	}

	return topic.Send(msg)
}

// Pop retrieves a message from Topic's internal queue.
func (m *Manager) Pop(name string) (*message.Message, error) {
	zap.L().Sugar().Infof("pop from '%s'", name)
	topic, ok := m.topics[name]
	if !ok {
		return nil, fmt.Errorf(errTopicNotFound, name)
	}

	return topic.Pop()
}
