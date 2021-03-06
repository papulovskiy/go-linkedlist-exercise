package linkedlist

import ()

type Item struct {
	Value interface{}
	next  *Item
}

type LinkedList struct {
	Head *Item
	Tail *Item
}

func (l *LinkedList) Insert(v interface{}) {
	item := &Item{v, nil}
	if l.Tail != nil {
		l.Tail.next = item
	}
	if l.Head == nil {
		l.Head = item
	}
	l.Tail = item
}

func (l *LinkedList) Contains(v interface{}) bool {
	cursor := l.Head
	if cursor == nil {
		return false
	}
	item := *cursor
	for {
		if item.Value == v {
			return true
		}
		if item.next == nil {
			break
		}
		item = *item.next
	}
	return false
}

func (l *LinkedList) FromArray(array []interface{}) {
	for _, item := range array {
		l.Insert(item)
	}
}

func FromArray(array []interface{}) LinkedList {
	var ll LinkedList
	ll.FromArray(array)
	return ll
}

func (l *LinkedList) Size() int {
	cursor := l.Head
	if cursor == nil {
		return 0
	}
	size := 1
	item := *cursor
	for {
		if item.next == nil {
			break
		}
		item = *item.next
		size++
	}
	return size
}
