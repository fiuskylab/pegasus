package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	t.Run("Validate", func(t *testing.T) {
		t.Run("Empty Body", func(t *testing.T) {
			require := require.New(t)
			msg := Message{
				TopicName: "something",
			}

			actual := msg.Validate()

			require.NotNil(actual)
		})

		t.Run("Empty TopicName", func(t *testing.T) {
			require := require.New(t)
			msg := Message{
				Body: "something",
			}

			actual := msg.Validate()

			require.NotNil(actual)
		})

		t.Run("Empty TopicName", func(t *testing.T) {
			require := require.New(t)
			msg := Message{
				Body:      "something",
				TopicName: "something",
			}

			actual := msg.Validate()

			require.Nil(actual)
		})
	})
}
