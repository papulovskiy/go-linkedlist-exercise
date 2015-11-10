package linkedlist_test

import (
	"./linkedlist"
	"testing"
)

func TestInsert(t *testing.T) {
	var ll linkedlist.LinkedList
	ll.Insert(1)
	ll.Insert("2")
}

func TestContains(t *testing.T) {
	var ll linkedlist.LinkedList

	if ll.Contains(1) {
		t.Fatal("List should be empty")
	}

	ll.Insert(1)
	if !ll.Contains(1) {
		t.Fatal("List should contain element")
	}

	ll.Insert("2")
	ll.Insert(3)
	if !ll.Contains(1) {
		t.Fatal("List should contain element")
	}
	if !ll.Contains("2") {
		t.Fatal("List should contain element")
	}
	if !ll.Contains(3) {
		t.Fatal("List should contain element")
	}
	if ll.Contains(4) {
		t.Fatal("List should not contain element")
	}

}
