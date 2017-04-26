package teaser

// Teaser:
// Design a queuing system. The queue should implement add, view and delete. A
// viewed message will be invisible to other workers for 1 second unless it is
// deleted. Messages should be returned in order they were added unless they have
// been deleted.
//
// The add method takes a string as a message and returns a unique id for that
// message.  The view method takes no parameters and returns a hash containing the
// next message and the unique message id assigned in the add method.  The delete
// method takes the unique message id and returns true if the message was removed
// within 1 second or false if we were too slow.
//
// Extra points: Do you see any problems with running this kind of queue in a
// production envrionment?
type queue struct{}

// New returns a new instance of the queuing system.
func New() queue {
	return queue{}
}

// Add adds a message to the queue and returns
// a unique id for the message
func (q queue) Add(msg string) string {
	return "uuid"
}

// View returns a message in FIFO order and a hash of the next message in the queue.
// Returned message will be invisible for 1 second for others.
func (q queue) View() (hash string, uuid string) {
	return "", "uuid"
}

// Delete deletes a message from the queue. Returns true if the message was
// removed within 1 second or false if we were too slow.
func (q queue) Delete(uuid string) bool {
	return false
}
