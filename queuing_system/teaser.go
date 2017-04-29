package teaser

import (
	"crypto/sha256"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

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
// TODO: get the extra points by providing an answer
type queue struct {
	head *node
	tail *node
}

// New returns a new instance of the queuing system.
func New() *queue {
	return &queue{head: sentinel(), tail: sentinel()}
}

// Add adds a message to the queue and returns
// a unique id (uuidv4) for the message.
//
// note:
//  passing the same message yields two different ids
//  messages of size greater than 1024 will be rejected
func (q *queue) Add(msg string) string {
	if len(msg) > 1024 {
		return ""
	}
	//TODO: add locking
	uid := uuid.NewV4()

	oldtail := q.tail
	q.tail = &node{data: msg, id: uid.String(), next: sentinel()}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldtail.next = q.tail
	}

	return uid.String()
}

// View returns a message in FIFO order and a sha-256 of the next message in the queue.
// Returned message will be invisible for 1 second for others.
//
// note this method uses a timestamp of type time.Now to determine if a message can be seen.
// the timestamp should not be considered monotonic, time can drift.
func (q *queue) View() (hash string, uuid string) {
	//TODO: add locking
	if q.isEmpty() {
		return "", ""
	}

	n := q.head
	for {
		if n.isSentinel() {
			break
		}
		if n.canBeSeen() {
			n.lastSeen = time.Now()
			return n.next.hash(), n.id
		}
		n = n.next
	}
	return n.hash(), n.id
}

// Delete deletes a message from the queue. Returns true if the message was
// removed within 1 second or false if we were too slow.
func (q *queue) Delete(uuid string) bool {
	// TODO: add locking
	if q.isEmpty() {
		return false
	}
	// TODO: input sanitization
	return false
}

// isEmpty determines whether the queue is empty
func (q *queue) isEmpty() bool {
	return q.head.isSentinel()
}

// node is a helper type used by the queue
// provides linked list semantics
type node struct {
	data     string
	id       string
	next     *node
	atype    string
	lastSeen time.Time
}

// isSentinel determines whether a node is of type sentinel
func (n node) isSentinel() bool {
	return n.atype == "sentinel"
}

// canBeSeen determines if a message can be seen by a caller.
// a viewed message is invisible for 1 second.
//
// note due to floating-point arithmeitc don't expect accuracy
// exactly 1 second.
func (n node) canBeSeen() bool {
	if n.lastSeen.IsZero() {
		return true
	}
	elapsed := time.Since(n.lastSeen)
	return elapsed.Seconds() > float64(1)
}

// hash returns sha-256 calculated over a message
func (n node) hash() string {
	if n.isSentinel() {
		return ""
	}
	h := sha256.New()
	h.Write([]byte(n.data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// sentinel creates a node of type sentinel
// think path terminator
func sentinel() *node {
	return &node{atype: "sentinel", next: &node{atype: "sentinel"}}
}
