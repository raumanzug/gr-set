package simple

import (
	"iter"
	"maps"

	"github.com/raumanzug/gr-set/set"
)

type simpleSet[T comparable] map[T]interface{}

// NewSet generates an empty set.
func NewSet[T comparable]() set.ISet[T] {
	return simpleSet[T](make(map[T]interface{}))
}

func (s simpleSet[T]) Clone() set.ISet[T] {
	return simpleSet[T](maps.Clone(s))
}

func (s simpleSet[T]) Generator() iter.Seq[T] {
	return maps.Keys(s)
}

func (s simpleSet[T]) Contains(elem T) bool {
	_, retval := s[elem]
	return retval
}

func (s simpleSet[T]) Add(elem T) {
	s[elem] = nil
}

func (s simpleSet[T]) AddSet(other set.ISet[T]) {
	pG := other.Generator()
	for elem := range pG {
		s.Add(elem)
	}
}

func (s simpleSet[T]) Remove(elem T) {
	delete(s, elem)
}

func (s simpleSet[T]) RemoveSet(other set.ISet[T]) {
	pG := other.Generator()
	for elem := range pG {
		s.Remove(elem)
	}
}

func (s simpleSet[T]) Retain(other set.ISet[T]) {
	for elem := range maps.Keys(s) {
		if !other.Contains(elem) {
			delete(s, elem)
		}
	}
}

func (s simpleSet[T]) Subeq(other set.ISet[T]) bool {

	for elem, _ := range s {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (s simpleSet[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s simpleSet[T]) Card() uint {
	return uint(len(s))
}
