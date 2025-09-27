package set

import (
	"iter"
)

// ISet is set type.
type ISet[T any] interface {

	// Clone produces another set containing the same elements.
	Clone() ISet[T]

	// Generator produces a generator which enumerates each element.
	Generator() iter.Seq[T]

	// Add adds an element elem to the set.
	Add(elem T)

	// Add adds each element of other to the set.
	AddSet(other ISet[T])

	// Remove removes element elem in set if set contains it.
	Remove(elem T)

	// Remove removes each element contained in other from set.
	RemoveSet(other ISet[T])

	// Retain removes each element not contained in other set.
	Retain(other ISet[T])

	// Contains checks whether set contains elem.
	Contains(elem T) bool

	// Subeq checks whether set is subset or equal to other.
	Subeq(other ISet[T]) bool

	// IsEmpty checks whether set is empty.
	IsEmpty() bool

	// Card yields the cardinality of set, i.e. the number of elements
	// the set contains.
	Card() uint
}
