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
//
//  Issue 1: unable to view a content of a message. The View method returns
//  an id of the message and the hash of a content of the next message. How can one read
//  a content of a message in this case?
//
//  Issue 2: you cannot be sure if a message was deleted or not. The Delete method returns
//  boolean value so you really don't know whether something went wrong
//  or perhaps you didn't fit in 1 second time window.
//
//  Issue 3: there is no guarantee that a message will not be read by others.
//  Given that one will call View() followed by immediate Delete().
//  network delays and packet reordering are not your friends.
type Queue struct {
	head *node
	tail *node
}

// New returns a new instance of the queuing system.
func New() *Queue {
	return &Queue{head: sentinel(), tail: sentinel()}
}

// Add adds a message to the queue and returns
// a unique id (uuidv4) for the message.
//
// note:
//  passing the same message yields two different ids
//  messages of size greater than 1024 will be rejected
//
// this method is considered thread unsafe
func (q *Queue) Add(msg string) string {
	if len(msg) > 1024 || len(msg) == 0 {
		return ""
	}

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

// View returns a message id in FIFO order and a sha-256 of the next message in the queue.
// Returned message will be invisible for 1 second for others.
//
// note this method uses a timestamp of type time.Now to determine if a message can be seen.
// the timestamp should not be considered monotonic, time can drift.
//
// this method is considered thread unsafe
func (q *Queue) View() (hash string, id string) {
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
//
// this method is considered thread unsafe
func (q *Queue) Delete(id string) bool {
	if q.isEmpty() {
		return false
	}
	_, err := uuid.FromString(id)
	if err != nil {
		return false
	}

	n := q.head
	found := false
	elapsed := time.Now()
	for {
		if n.isSentinel() {
			break
		}
		if n.id == id {
			q.head = n.next
			found = true
			break
		} else if n.next.id == id {
			n.next = n.next.next
			found = true
			break
		}
		n = n.next
	}
	return found && time.Since(elapsed).Seconds() < float64(1)
}

// isEmpty determines whether the queue is empty
func (q *Queue) isEmpty() bool {
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
