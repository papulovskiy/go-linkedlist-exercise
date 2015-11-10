package linkedlist_test

import (
	"./linkedlist"
	"math/rand"
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

func TestArray(t *testing.T) {
	var ll linkedlist.LinkedList

	array := make([]interface{}, 7)
	array[0] = "first string"
	array[1] = "just a string"
	for i := 2; i < 7; i++ {
		array[i] = rand.Int()
	}

	ll.FromArray(array)
	for i := 0; i < 7; i++ {
		if !ll.Contains(array[i]) {
			t.Fatal("List should contain an array item")
		}
	}
	if ll.Contains("just another string") {
		t.Fatal("List should not contain item not from array")
	}
}
