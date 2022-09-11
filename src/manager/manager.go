package manager

import (
	"fmt"

	"github.com/fiuskylab/pegasus/src/topic"
)

type (
	// Manager - Manages Pegasus topics.
	Manager struct {
		topics map[string]*topic.Topic
	}
)

// NewManager returns a new Manager
func NewManager() *Manager {
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
)

// NewTopic creates a new Topic and adds it to the Manager.
func (m *Manager) NewTopic(name string) error {
	if _, ok := m.topics[name]; ok {
		return fmt.Errorf(errTopicAlreadyExists, name)
	}
	createdTopic := topic.NewTopic(name)

	m.topics[createdTopic.Name] = createdTopic
	return nil
}
