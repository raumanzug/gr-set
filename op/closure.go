package op

import (
	"iter"

	"github.com/raumanzug/gr-set/set"
	"github.com/raumanzug/gr-set/simple"
)

// Closure closes base under applying op.
//
// i.e. after performing Closure base is the least set containing
// each element of base before performing Closure and if
// op(x) =  [y_1, ... y_m] with m >= 0 and x is contained in base
// then also y_i for each i in {1, ..., m} is contained in base as well.
func Closure[T comparable](
	base set.ISet[T],
	op func(T) iter.Seq[T]) {

	bagA := base.Generator()
	for {
		bagB := simple.NewSet[T]()
		isEmpty := true
		for elem := range bagA {
			isEmpty = false
			for next := range op(elem) {
				if !base.Contains(next) {
					bagB.Add(next)
				}
			}
		}
		if isEmpty {
			break
		}
		base.AddSet(bagB)
		bagA = bagB.Generator()
	}
}
