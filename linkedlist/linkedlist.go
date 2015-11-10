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
			return false
		}
		item = *item.next
	}
	return false
}
