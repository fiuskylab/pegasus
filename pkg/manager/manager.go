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
