package comparable

import (
	"iter"
	"slices"
	"sort"

	"github.com/raumanzug/gr-set/set"
)

type comparableSet_t[C any, T set.IComparable[C, T]] struct {
	ctxt C
	data []T
}

func (recv comparableSet_t[C, T]) Len() int {
	return len(recv.data)
}

func (recv comparableSet_t[C, T]) Less(i, j int) bool {
	return recv.data[i].CompareTo(recv.ctxt, recv.data[j]) < 0
}

func (recv comparableSet_t[C, T]) Swap(i, j int) {
	recv.data[i], recv.data[j] = recv.data[j], recv.data[i]
}

// NewSet generates a new set consisting in the elements in argument elems.
//
// elems will be modified (sorted according to the order relation defined
// by interface set.IComparable).
func NewSet[C any, T set.IComparable[C, T]](
	ctxt C,
	elems ...T,
) set.ISortedSet[T] {
	retval := comparableSet_t[C, T]{
		ctxt: ctxt,
		data: elems,
	}

	sort.Sort(retval)
	retval.data = slices.CompactFunc(
		elems,
		func(x T, y T) bool {
			return x.CompareTo(ctxt, y) == 0
		},
	)

	return &retval
}

func (recv comparableSet_t[C, T]) Clone() set.ISet[T] {
	return &comparableSet_t[C, T]{
		ctxt: recv.ctxt,
		data: slices.Clone(recv.data),
	}
}

func (recv comparableSet_t[C, T]) Generator() iter.Seq[T] {
	// clone it because generator will be confused if recv.data is exchanged
	// by Remove and other methods.
	return slices.Values(slices.Clone(recv.data))
}

func (recv comparableSet_t[C, T]) Contains(elem T) bool {
	_, isFound := sort.Find(
		len(recv.data),
		func(index int) int {
			return elem.CompareTo(recv.ctxt, recv.data[index])
		},
	)

	return isFound
}

func (pRecv *comparableSet_t[C, T]) Add(elem T) {
	index, isFound := sort.Find(
		len(pRecv.data),
		func(index int) int {
			return elem.CompareTo(pRecv.ctxt, pRecv.data[index])
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

func (pRecv *comparableSet_t[C, T]) AddSet(other set.ISet[T]) {
	pG := other.Generator()
	for elem := range pG {
		pRecv.Add(elem)
	}
}

func (pRecv *comparableSet_t[C, T]) Remove(elem T) {
	index, isFound := sort.Find(
		len(pRecv.data),
		func(index int) int {
			return elem.CompareTo(pRecv.ctxt, pRecv.data[index])
		},
	)

	if !isFound {
		return
	}

	pRecv.data = slices.Delete(pRecv.data, index, index+1)
}

func (pRecv *comparableSet_t[C, T]) RemoveSet(other set.ISet[T]) {
	pG := other.Generator()
	for elem := range pG {
		pRecv.Remove(elem)
	}
}

func (pRecv *comparableSet_t[C, T]) Representative(elem T) T {
	index, isFound := sort.Find(
		len(pRecv.data),
		func(index int) int {
			return elem.CompareTo(pRecv.ctxt, pRecv.data[index])
		},
	)

	if isFound {
		return pRecv.data[index]
	} else {
		pRecv.Add(elem)
		return elem
	}
}

func (pRecv *comparableSet_t[C, T]) Retain(other set.ISet[T]) {
	for _, elem := range pRecv.data {
		if !other.Contains(elem) {
			pRecv.Remove(elem)
		}
	}
}

func (pRecv *comparableSet_t[C, T]) Sort() {
	sort.Sort(pRecv)
	pRecv.data = slices.CompactFunc(
		pRecv.data,
		func(x T, y T) bool {
			return x.CompareTo(pRecv.ctxt, y) == 0
		},
	)
}

func (recv comparableSet_t[C, T]) Subeq(other set.ISet[T]) bool {

	for _, elem := range recv.data {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (pRecv *comparableSet_t[C, T]) Eq(other set.ISet[T]) bool {
	return pRecv.Subeq(other) && other.Subeq(pRecv)
}

func (recv comparableSet_t[C, T]) IsDisjoint(other set.ISet[T]) bool {

	for _, elem := range recv.data {
		if other.Contains(elem) {
			return false
		}
	}

	return true
}

func (recv comparableSet_t[C, T]) IsEmpty() bool {
	return len(recv.data) == 0
}

func (recv comparableSet_t[C, T]) Card() uint {
	return uint(len(recv.data))
}
