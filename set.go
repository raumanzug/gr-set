package set

import (
	"iter"
)

// Set is set type.
type Set[T any] interface {

	// Clone produces another set containing the same elements.
	Clone() Set[T]

	// Generator produces a generator which enumerates each element.
	Generator() iter.Seq[T]

	// Add adds an element elem to the set.
	Add(elem T)

	// Add adds each element of other to the set.
	AddSet(other Set[T])

	// Remove removes element elem in set if set contains it.
	Remove(elem T)

	// Remove removes each element contained in other from set.
	RemoveSet(other Set[T])

	// Retain removes each element not contained in other from set.
	Retain(other Set[T])

	// Contains checks whether set contains elem.
	Contains(elem T) bool

	// Subeq checks whether set is subset or equal to other.
	Subeq(other Set[T]) bool

	// IsEmpty checks whether set is empty.
	IsEmpty() bool

	// Card yields the cardinality of set, i.e. the number of elements
	// the set contains.
	Card() uint
}
