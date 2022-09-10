package message

type (
	// Message - A data model that represents the messages
	// sent/received to/from services subscribed to a Topic.
	Message struct {
		Body string `json:"body"`
		// TODO:
		// 	- Add more features, like:
		// 		- Attributes
		// 		- Expiration
		// 		- ID ???
	}
)
