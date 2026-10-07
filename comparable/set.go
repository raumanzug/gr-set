package comparable

import (
	"iter"
	"slices"
	"sort"

	"github.com/raumanzug/gr-set/set"
)

type comparableSet_t[T set.IComparable[T]] struct {
	data []T
}

func (recv comparableSet_t[T]) Len() int {
	return len(recv.data)
}

func (recv comparableSet_t[T]) Less(i, j int) bool {
	return recv.data[i].CompareTo(recv.data[j]) < 0
}

func (recv comparableSet_t[T]) Swap(i, j int) {
	recv.data[i], recv.data[j] = recv.data[j], recv.data[i]
}

// NewSet generates a new set consisting in the elements in argument elems.
//
// elems will be modified (sorted according to the order relation defined
// by interface set.IComparable).
func NewSet[T set.IComparable[T]](elems ...T) set.ISet[T] {
	retval := comparableSet_t[T]{
		data: elems,
	}

	sort.Sort(retval)
	retval.data = slices.CompactFunc(
		elems,
		func(x T, y T) bool {
			return x.CompareTo(y) == 0
		},
	)

	return &retval
}

func (recv comparableSet_t[T]) Clone() set.ISet[T] {
	return &comparableSet_t[T]{
		data: slices.Clone(recv.data),
	}
}

func (recv comparableSet_t[T]) Generator() iter.Seq[T] {
	// clone it because generator will be confused if recv.data is exchanged
	// by Remove and other methods.
	return slices.Values(slices.Clone(recv.data))
}

func (recv comparableSet_t[T]) Contains(elem T) bool {
	_, isFound := sort.Find(
		len(recv.data),
		func(index int) int {
			return elem.CompareTo(recv.data[index])
		},
	)

	return isFound
}

func (pRecv *comparableSet_t[T]) Add(elem T) {
	index, isFound := sort.Find(
		len(pRecv.data),
		func(index int) int {
			return elem.CompareTo(pRecv.data[index])
		},
	)

	if isFound {
		return
	}

	if index < len(pRecv.data) {
		pRecv.data = slices.Insert(
			pRecv.data,
			index,
			elem,
		)
	} else {
		pRecv.data = append(pRecv.data, elem)
	}
}

func (pRecv *comparableSet_t[T]) AddSet(other set.ISet[T]) {
	pG := other.Generator()
	for elem := range pG {
		pRecv.Add(elem)
	}
}

func (pRecv *comparableSet_t[T]) Remove(elem T) {
	index, isFound := sort.Find(
		len(pRecv.data),
		func(index int) int {
			return elem.CompareTo(pRecv.data[index])
		},
	)

	if !isFound {
		return
	}

	pRecv.data = slices.Delete(pRecv.data, index, index+1)
}

func (pRecv *comparableSet_t[T]) RemoveSet(other set.ISet[T]) {
	pG := other.Generator()
	for elem := range pG {
		pRecv.Remove(elem)
	}
}

func (pRecv *comparableSet_t[T]) Representative(elem T) T {
	index, isFound := sort.Find(
		len(pRecv.data),
		func(index int) int {
			return elem.CompareTo(pRecv.data[index])
		},
	)

	if isFound {
		return pRecv.data[index]
	} else {
		pRecv.Add(elem)
		return elem
	}
}

func (pRecv *comparableSet_t[T]) Retain(other set.ISet[T]) {
	for _, elem := range pRecv.data {
		if !other.Contains(elem) {
			pRecv.Remove(elem)
		}
	}
}

func (recv comparableSet_t[T]) Subeq(other set.ISet[T]) bool {

	for _, elem := range recv.data {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (pRecv *comparableSet_t[T]) Eq(other set.ISet[T]) bool {
	return pRecv.Subeq(other) && other.Subeq(pRecv)
}

func (recv comparableSet_t[T]) IsDisjoint(other set.ISet[T]) bool {

	for _, elem := range recv.data {
		if other.Contains(elem) {
			return false
		}
	}

	return true
}

func (recv comparableSet_t[T]) IsEmpty() bool {
	return len(recv.data) == 0
}

func (recv comparableSet_t[T]) Card() uint {
	return uint(len(recv.data))
}
