package set

import (
	"iter"
)

// ISet is set type.
type ISet[T any] interface {

	// Add adds an element elem to the set.
	Add(elem T)

	// Add adds each element of other to the set.
	AddSet(other ISet[T])

	// Card yields the cardinality of set, i.e. the number of elements
	// the set contains.
	Card() uint

	// Clone produces another set containing the same elements.
	Clone() ISet[T]

	// Contains checks whether set contains elem.
	Contains(elem T) bool

	// Eq checks whether set equals to other.
	Eq(other ISet[T]) bool

	// Generator produces a generator which enumerates each element.
	Generator() iter.Seq[T]

	// IsDisjoint checks whether other set is disjoint to receiver set.
	IsDisjoint(other ISet[T]) bool

	// IsEmpty checks whether set is empty.
	IsEmpty() bool

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

	// Subeq checks whether set is subset or equal to other.
	Subeq(other ISet[T]) bool
}

// ISetMap is used for looking up elements using ids.
type ISetMap[I any, S any, T any] interface {

	// Get looks up element with id `id` and returns element, its
	// index and true if an element exists with this id.
	// If not, it returns false.
	Get(id S) (I, T, bool)

	// Remove removes element with index `index`.
	Remove(index I)
}

// ISortedSet is the interface for implementations of the ISet interface
// that require a strict ordering of their elements.
type ISortedSet[T any] interface {
	ISet[T]

	// Sort is necessary after modifying elements managed by
	// comparableSet_t structures.  Otherwise the slice inside this
	// structure turns to be not ordered according to ordering defined by
	// IComparable.
	Sort()
}
