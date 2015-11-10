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
	l.Tail = Item{v, nil}
}

func (l *LinkedList) Contains() bool {

}
