package teaser_test

import (
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
		t.Fatal("added only one msg but the hash of the next one returned")
	}

	deleted := target.Delete(mid)
	if !deleted {
		t.Fatalf("the message = %s was not deleted", mid)
	}
}
