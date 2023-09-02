package set

import (
	"github.com/raumanzug/gr-generator"
)

// Closure closes base under applying op.
//
// i.e. after performing Closure base is the least set containing
// each element of base before performing Closure and if
// op(x) =  [y_1, ... y_m] with m >= 0 and x is contained in base
// then also y_i for each i in {1, ..., m} is contained in base as well.
func Closure[T comparable](
	base Set[T],
	op func(T) generator.Generator[generator.LoopDirective, T]) {

	bagA := base.Generator()
	for {
		bagB := NewSimpleSet[T]()
		isEmpty := true
		generator.Foreach(
			bagA,
			func(elem T) generator.LoopDirective {
				isEmpty = false
				generator.Foreach(
					op(elem),
					func(next T) generator.LoopDirective {
						if !base.Contains(next) {
							bagB.Add(next)
						}
						return generator.LoopDirectiveContinue
					})
				return generator.LoopDirectiveContinue
			})
		if isEmpty {
			break
		}
		base.AddSet(bagB)
		bagA = bagB.Generator()
	}
}
