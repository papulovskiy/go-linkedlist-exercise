package set_test

import (
	"github.com/papulovskiy/go-linkedlist-exercise/set"
	// "math/rand"
	"testing"
)

func TestInsertAndContains(t *testing.T) {
	var s set.Set

	if s.Contains(1) {
		t.Fatal("Set should be empty")
	}

	s.Add(1)
	if !s.Contains(1) {
		t.Fatal("Set should contain element")
	}
}

func TestUniqueness(t *testing.T) {
	var s set.Set
	if s.Size != 0 {
		t.Fatal("Set is not empty")
	}

	s.Add(1)
	if s.Size != 1 {
		t.Fatal("Set size is not correct")
	}

	s.Add(2)
	if s.Size != 2 {
		t.Fatal("Set size is not correct")
	}

	s.Add(1)
	if s.Size != 2 {
		t.Fatal("Set size is not correct")
	}
}
