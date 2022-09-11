package topic

import (
	"regexp"
	"testing"

	"github.com/fiuskylab/pegasus/src/message"
	"github.com/stretchr/testify/require"
)

var (
	uuidRegex = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89a][0-9a-f]{3}-[0-9a-f]{12}$")
)

func TestTopic(t *testing.T) {
	t.Run("NewTopic", func(t *testing.T) {
		t.Run("With name", func(t *testing.T) {
			require := require.New(t)
			name := "topic"

			actual, err := NewTopic(name)

			require.Nil(err)
			require.NotNil(actual)
			require.Equal(name, actual.Name)
		})
		t.Run("Empty name", func(t *testing.T) {
			require := require.New(t)

			actual, err := NewTopic("")

			require.Nil(err)
			require.NotNil(actual)
			require.Regexp(uuidRegex, actual.Name)
		})
	})

	t.Run("Send", func(t *testing.T) {
		require := require.New(t)
		topic, err := NewTopic("")

		require.Nil(err)

		actual := topic.Send(&message.Message{
			Body: "something",
		})

		require.Nil(actual)
	})

	t.Run("Pop", func(t *testing.T) {
		require := require.New(t)
		topic, err := NewTopic("")
		msg := message.Message{
			Body: "somethin",
		}

		require.Nil(err)

		{
			err := topic.Send(&msg)
			require.Nil(err)
		}

		{
			actual, err := topic.Pop()
			require.Nil(err)
			require.Equal(msg, *actual)
		}
	})
}
