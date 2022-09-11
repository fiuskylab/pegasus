package manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestManager(t *testing.T) {
	t.Run("NewManager", func(t *testing.T) {
		require := require.New(t)

		actual := NewManager()

		require.NotNil(actual)
	})

	t.Run("NewTopic", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()
			actualErr := m.NewTopic("")

			require.Nil(actualErr)
		})

		t.Run("Topic Exists", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()

			{
				actual := m.NewTopic("topic")
				require.Nil(actual)
			}
			{
				actual := m.NewTopic("topic")
				require.NotNil(actual)
			}
		})
	})

	t.Run("GetTopicNames", func(t *testing.T) {
		t.Run("Success - Empty", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()

			actual := m.GetTopicNames()

			require.Empty(actual)
		})

		t.Run("Success - 1 Topic", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()

			{
				actual := m.NewTopic("topic")
				require.Nil(actual)
			}

			{
				expected := []string{"topic"}
				actual := m.GetTopicNames()

				require.Equal(expected, actual)
			}
		})

		t.Run("Success - 10 Topics", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()

			amountOfTopics := 10

			{
				for i := 0; i < amountOfTopics; i++ {
					actual := m.NewTopic("")
					require.Nil(actual)
				}
			}

			{
				topics := m.GetTopicNames()

				actual := len(topics)
				require.Equal(amountOfTopics, actual)
			}
		})
	})
}
