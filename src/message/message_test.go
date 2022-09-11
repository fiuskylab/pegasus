package message

import (
	"testing"

	"github.com/fiuskylab/pegasus/src/proto"
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

	t.Run("FromRequest", func(t *testing.T) {
		t.Run("Nil Request", func(t *testing.T) {
			require := require.New(t)

			_, actual := FromRequest(nil)

			require.NotNil(actual)
		})

		t.Run("Empty Body", func(t *testing.T) {
			require := require.New(t)

			_, actual := FromRequest(&proto.SendRequest{
				Body:      "",
				TopicName: "topic",
			})

			require.NotNil(actual)
		})

		t.Run("Empty TopicName", func(t *testing.T) {
			require := require.New(t)

			_, actual := FromRequest(&proto.SendRequest{
				Body:      "body",
				TopicName: "",
			})

			require.NotNil(actual)
		})

		t.Run("Valid", func(t *testing.T) {
			require := require.New(t)

			expected := Message{
				Body:      "body",
				TopicName: "topic",
			}

			actual, err := FromRequest(&proto.SendRequest{
				Body:      "body",
				TopicName: "topic",
			})

			require.Nil(err)
			require.Equal(expected, *actual)
		})
	})
}
