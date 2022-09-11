# Notes

### What is a "Message-broker"?
_"A message broker (also known as an integration broker or interface engine) is an intermediary computer program module that translates a message from the formal messaging protocol of the sender to the formal messaging protocol of the receiver. [...]"_

### What is Pegasus?
Pegasus is a open-source minimum message broker written in Go

### Key Features
- Topic
  - This is a "named queue", where services can subscribe to.
- Messages
  - Those are items sent/received to/from _Topics_.
  - Can have an expiration
- Storage
  - We need to store the messages/topics somewhere.

### Communication
1. gRPC
  - Can subscribe to a bidirectional streaming "endpoint"
  - Need to generate the _proto_ files for a given language (CLI ??)
2. Streaming HTTP/2
  - Can subscribe to a bidirectional streaming "endpoint"
  - Client package for each language or a Swagger doc.
3. Pegasus Protocol
  - TODO!
  - Client package for each language
