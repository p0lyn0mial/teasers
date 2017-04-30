package teaser_test

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	teaser "github.com/teasers/queuing_system"
)

// TestAddViewDelete simple test case that will
// check single add, view, delete
func TestAddViewDelete(t *testing.T) {
	target := teaser.New()
	mid := target.Add("hello_world")

	{
		nextHash, mvid := target.View()
		if mid != mvid {
			t.Fatalf("mismatch of added msg id = %s, and viewed msg id = %s", mid, mvid)
		}
		if len(nextHash) != 0 {
			t.Fatalf("added only one msg but the hash of the next one returned = %s", nextHash)
		}

		deleted := target.Delete(mid)
		assertDeleted(mid, deleted, t)
	}

	{
		// check if we neither get nor delete a message
		nextHash, nextvid := target.View()
		assertMessageId("", nextvid, t)
		assertMessageHash("", nextHash, t)

		deleted := target.Delete(mid)
		if deleted {
			assertDeleted(mid, deleted, t)
		}
	}
}

// TestAddTwoDeleteFirst checks if the first msg is deleted, the second
// one is viewable
func TestAddTwoDeleteFirst(t *testing.T) {
	target := teaser.New()
	m1id := target.Add("msg1")
	m2id := target.Add("msg2")

	{
		deleted := target.Delete(m1id)
		assertDeleted(m1id, deleted, t)
	}

	{
		nextHash, m2vid := target.View()
		assertMessageId(m2id, m2vid, t)
		assertMessageHash("", nextHash, t)
	}
}

// TestAddTwoDeleteSecond checks if the second msg is deleted,
// the first one is viewable
func TestAddTwoDeleteScond(t *testing.T) {
	target := teaser.New()
	m1id := target.Add("msg1")
	m2id := target.Add("msg2")

	{
		deleted := target.Delete(m2id)
		assertDeleted(m2id, deleted, t)
	}

	{
		nextHash, m1vid := target.View()
		assertMessageId(m1id, m1vid, t)
		assertMessageHash("", nextHash, t)
	}
}

// TestAddManyRemoveFew checks if removing messages from the middle of the queue works.
func TestAddManyRemoveFew(t *testing.T) {
	target := teaser.New()
	mids := []string{}
	for i := 0; i <= 10; i++ {
		mid := target.Add(fmt.Sprintf("msg%d", i))
		mids = append(mids, mid)
	}
	if len(mids) < 10 {
		t.Fatal("broken test data, expected mids to have at least 9 elementes ")
	}

	{
		m5id := mids[5]
		deleted := target.Delete(m5id)
		assertDeleted(m5id, deleted, t)
		assertIsNotInTheQueue(target, m5id, t)

		newmids := mids[:5]
		newmids = append(newmids, mids[6:]...)
		mids = newmids
		waitForMessages()
		assertTheQueueContainsIds(target, mids, t)
		waitForMessages()
	}

	{
		m5id := mids[5]
		deleted := target.Delete(m5id)
		assertDeleted(m5id, deleted, t)
		assertIsNotInTheQueue(target, m5id, t)
		nmids := mids[:5]
		nmids = append(nmids, mids[6:]...)
		mids = nmids
		waitForMessages()
		assertTheQueueContainsIds(target, mids, t)
	}
}

// TestAddOnceViewMany checks if subsequent calls return an empty set
func TestAddOnceViewMany(t *testing.T) {
	target := teaser.New()
	m1id := target.Add("msg1")

	{
		nextHash, m1vid := target.View()
		assertMessageId(m1id, m1vid, t)
		assertMessageHash("", nextHash, t)
	}

	{
		nextHash, nextvid := target.View()
		assertMessageId("", nextvid, t)
		assertMessageHash("", nextHash, t)
	}
}

// TestAddOnceViewManySleep checks if a message can be seen
// after 1 s
func TestAddOnceViewManySleep(t *testing.T) {
	target := teaser.New()
	m1id := target.Add("msg1")

	{
		nextHash, m1vid := target.View()
		assertMessageId(m1id, m1vid, t)
		assertMessageHash("", nextHash, t)
	}

	{
		nextHash, nextvid := target.View()
		assertMessageId("", nextvid, t)
		assertMessageHash("", nextHash, t)
	}

	{
		waitForMessages()
		nextHash, m1vid := target.View()
		assertMessageId(m1id, m1vid, t)
		assertMessageHash("", nextHash, t)
	}
}

// TestAddManyViewMany checks if subseqnet calls return valid data
func TestAddManyViewMany(t *testing.T) {
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

func assertDeleted(mid string, deleted bool, t *testing.T) {
	if !deleted {
		t.Fatalf("the message = %s was not deleted", mid)
	}
}

func waitForMessages() {
	// yeah I know sleeping is not the best synchronization primitive
	time.Sleep(time.Duration(2) * time.Second)
}

func assertIsNotInTheQueue(target *teaser.Queue, id string, t *testing.T) {
	for {
		_, mid := target.View()
		if len(mid) == 0 {
			break
		}
		if mid == id {
			t.Fatalf("the message = %s was found in the queue", id)
		}
	}
}

func assertTheQueueContainsIds(target *teaser.Queue, ids []string, t *testing.T) {
	found := []string{}
	for {
		_, mid := target.View()
		if len(mid) == 0 {
			break
		}
		found = append(found, mid)
	}
	if len(found) != len(ids) {
		t.Fatalf("found unexpected number of ids = %d, expected to find %d", len(found), len(ids))
	}
	for index, id := range ids {
		if found[index] != id {
			t.Fatalf("incorrect id found at index = %d, have = %s, expected = %s", index, found[index], id)
		}
	}
}
