package set

import (
	"github.com/raumanzug/gr-generator"
)

type simpleSet[T comparable] struct {
	data map[T]interface{}
}

// NewSimpleSet generates an empty set.
func NewSimpleSet[T comparable]() *simpleSet[T] {
	return &simpleSet[T]{make(map[T]interface{})}
}

func (s *simpleSet[T]) Clone() Set[T] {
	retval := NewSimpleSet[T]()

	for elem, _ := range s.data {
		retval.Add(elem)
	}

	return retval
}

type simpleSetGenerator[T comparable] struct {
	generator.GeneratorBase[generator.LoopDirective, T]
	set *simpleSet[T]
}

func (ssg *simpleSetGenerator[T]) Loop() {
	for elem, _ := range ssg.set.data {
		if generator.LoopDirectiveBreak == ssg.Yield(elem) {
			break
		}
	}
}

func (s *simpleSet[T]) Generator() generator.Generator[generator.LoopDirective, T] {
	return &simpleSetGenerator[T]{set: s}
}

func (s *simpleSet[T]) Contains(elem T) (retval bool) {
	_, retval = s.data[elem]
	return
}

func (s *simpleSet[T]) Add(elem T) {
	s.data[elem] = nil
}

func (s *simpleSet[T]) AddSet(other Set[T]) {
	pG := other.Generator()
	action := func(elem T) generator.LoopDirective {
		s.Add(elem)
		return generator.LoopDirectiveContinue
	}
	generator.Foreach(pG, action)
}

func (s *simpleSet[T]) Remove(elem T) {
	delete(s.data, elem)
}

func (s *simpleSet[T]) RemoveSet(other Set[T]) {
	pG := other.Generator()
	action := func(elem T) generator.LoopDirective {
		s.Remove(elem)
		return generator.LoopDirectiveContinue
	}
	generator.Foreach(pG, action)
}

func (s *simpleSet[T]) Retain(other Set[T]) {
	for elem, _ := range s.data {
		if !other.Contains(elem) {
			delete(s.data, elem)
		}
	}
}

func (s *simpleSet[T]) Subeq(other Set[T]) bool {

	for elem, _ := range s.data {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *simpleSet[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *simpleSet[T]) Card() uint {
	return uint(len(s.data))
}
