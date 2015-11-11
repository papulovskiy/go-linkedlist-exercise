package set

import (
	"github.com/papulovskiy/go-linkedlist-exercise/linkedlist"
)

type Set struct {
	list linkedlist.LinkedList
}

func (s *Set) Contains(v interface{}) bool {
	return s.list.Contains(v)
}

func (s *Set) Add(v interface{}) {
	if !s.Contains(v) {
		s.list.Insert(v)
	}
}

func (s *Set) Size() int {
	return s.list.Size()
}
