package set

import (
	"github.com/raumanzug/gr-generator"
	"testing"
)

func Test_IsEmpty(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	if !mySet.IsEmpty() {
		t.Fatalf("freshly initialized simple set should be empty.  is indeed not.")
	}
}

func Test_IsNonEmpty(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(12)
	if mySet.IsEmpty() {
		t.Fatalf("set after adding some elements should not be empty.  is indeed empty.")
	}
}

func Test_IsSubeqA(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setB := NewSimpleSet[uint]()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_IsSubeqB(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setA.Add(12)
	setB := setA.Clone()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_IsSubeqC(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setB := setA.Clone()
	setB.Add(12)
	if setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_AddIdempotenceA(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setB := NewSimpleSet[uint]()
	setA.Add(12)
	setB.Add(12)
	setB.Add(12)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_AddIdempotenceB(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setB := NewSimpleSet[uint]()
	setA.Add(12)
	setA.Add(12)
	setB.Add(12)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_AddCommutativity(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setB := NewSimpleSet[uint]()
	setA.Add(5)
	setA.Add(12)
	setB.Add(12)
	setB.Add(5)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_RemoveItself(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySet.RemoveSet(mySet)
	if !mySet.IsEmpty() {
		t.Fail()
	}
}

func Test_AddItselfA(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_AddItselfB(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_RetainItselfA(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_RetainItselfB(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_RetainA(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setA.Add(5)
	setB := setA.Clone()
	setA.Add(12)
	setA.Retain(setB)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_RetainB(t *testing.T) {
	setA := NewSimpleSet[uint]()
	setA.Add(5)
	setB := setA.Clone()
	setA.Add(12)
	setA.Retain(setB)
	if !setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_CardA(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySet.Add(5)
	if mySet.Card() != 2 {
		t.Fail()
	}
}

func Test_Break(t *testing.T) {
	mySet := NewSimpleSet[uint]()
	mySet.Add(37)
	mySet.Add(5)
	mySet.Add(12)
	g := mySet.Generator()
	resultList := []uint{}
	generator.Foreach(
		g,
		func(elem uint) generator.LoopDirective {
			if elem < 23 {
				return generator.LoopDirectiveBreak
			}
			resultList = append(resultList, elem)
			return generator.LoopDirectiveContinue
		})
	if len(resultList) > 2 {
		t.Fail()
	}
}
