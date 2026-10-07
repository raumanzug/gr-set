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

	// Add `elem` if set does not contain an equivalent for it.
	//
	//
	// If set does not contain an equivalent this method returns
	// `elem`.
	//
	// If set contains an equivalent this method returns this
	// equivalent and ignores `elem`.
	Representative(elem T) T

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

	// Eq checks whether set equals to other.
	Eq(other ISet[T]) bool

	// IsDisjoint checks whether other set is disjoint to receiver set.
	IsDisjoint(other ISet[T]) bool

	// IsEmpty checks whether set is empty.
	IsEmpty() bool

	// Card yields the cardinality of set, i.e. the number of elements
	// the set contains.
	Card() uint
}
