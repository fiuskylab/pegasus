package manager

import "github.com/fiuskylab/pegasus/pkg/topic"

type (
	// Manager - Manages Pegasus topics.
	Manager struct {
		topics map[string]topic.Topic
	}
)

// NewManager returns a new Manager
func NewManager() *Manager {
	return &Manager{
		topics: map[string]topic.Topic{},
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
