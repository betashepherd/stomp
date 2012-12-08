package queue

import (
	"github.com/jjeffery/stomp/message"
)

// Interface for queue storage. The intent is that
// different queue storage implementations can be
// used, depending on preference. Queue storage
// mechanisms could include in-memory, and various
// persistent storage mechanisms (eg file system, SQL, etc)
type QueueStorage interface {
	// Pushes a frame to the end of the queue. Sets
	// the "message-id" header of the frame before adding to
	// the queue.
	Enqueue(queue string, frame *message.Frame) error

	// Pushes a frame to the head of the queue. Sets
	// the "message-id" header of the frame if it is not
	// already set.
	Requeue(queue string, frame *message.Frame) error

	// Removes a frame from the head of the queue.
	// Returns nil if no frame is available.
	Dequeue(queue string) (*message.Frame, error)

	// Called prior to server shutdown. Allows the queue storage
	// to perform any cleanup.
	Stop()
}

type QueueManager struct {
	qstore QueueStorage
}

// Create a queue manager with the specified queue storage mechanism
func NewQueueManager(qstore QueueStorage) *QueueManager {
	qm := new(QueueManager)
	qm.qstore = qstore
	return qm
}