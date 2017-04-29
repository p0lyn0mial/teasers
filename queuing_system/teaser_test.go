package teaser_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	teaser "github.com/teasers/queuing_system"
)

// TestAddViewDelete simple test case that will
// check single add, view, delete
func TestAddViewDelete(t *testing.T) {
	target := teaser.New()
	mid := target.Add("hello_world")

	nextHash, mvid := target.View()
	if mid != mvid {
		t.Fatalf("mismatch of added msg id = %s, and viewed msg id = %s", mid, mvid)
	}
	if len(nextHash) != 0 {
		t.Fatalf("added only one msg but the hash of the next one returned = %s", nextHash)
	}

	deleted := target.Delete(mid)
	if !deleted {
		t.Fatalf("the message = %s was not deleted", mid)
	}
}

// TestAddViewMany a test case that adds many messages
// and then views them.
func TestAddViewMany(t *testing.T) {
	target := teaser.New()
	m1id := target.Add("msg1")
	m2id := target.Add("msg2")
	m3id := target.Add("msg3")

	m2Hash, m1vid := target.View()
	assertMessageId(m1id, m1vid, t)
	assertMessageHash("msg2", m2Hash, t)

	m3Hash, m2vid := target.View()
	assertMessageId(m2id, m2vid, t)
	assertMessageHash("msg3", m3Hash, t)

	nextHash, m3vid := target.View()
	assertMessageId(m3id, m3vid, t)
	assertMessageHash("", nextHash, t)

}

func assertMessageId(actual, expected string, t *testing.T) {
	if actual != expected {
		t.Fatalf("mismatch of added msg id = %s, and viewed msg id = %s", actual, expected)
	}
}

func assertMessageHash(msg, expectedHash string, t *testing.T) {
	h := sha256.New()
	h.Write([]byte(msg))
	actualHash := fmt.Sprintf("%x", h.Sum(nil))
	if len(msg) == 0 {
		actualHash = ""
	}
	if actualHash != expectedHash {
		t.Fatalf("mismatch of hashes over the msg = %s, actual = %s, expected = %s", msg, actualHash, expectedHash)
	}
}
