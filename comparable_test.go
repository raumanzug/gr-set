package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/comparable"
)

type testType_t struct {
	a uint
	b uint
}

func (recv testType_t) CompareTo(other testType_t) int {
	sum_left := recv.a + recv.b
	sum_right := other.a + other.b
	switch {
	case sum_left > sum_right:
		return 1
	case sum_left == sum_right:
		return 0
	case sum_left < sum_right:
		return -1
	}

	return 0
}

func Test_Comparable_IsEmpty(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	if !mySet.IsEmpty() {
		t.Fatalf("freshly initialized comparable set should be empty.  is indeed not.")
	}
}

func Test_Comparable_IsNonEmpty(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 7, b: 5})
	if mySet.IsEmpty() {
		t.Fatalf("set after adding some elements should not be empty.  is indeed empty.")
	}
}

func Test_Comparable_IsSubeqA(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsSubeqB(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB := setA.Clone()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsSubeqC(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := setA.Clone()
	setB.Add(testType_t{a: 7, b: 5})
	if setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_Comparable_AddIdempotenceA(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_AddIdempotenceB(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_AddCommutativity(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 2, b: 3})
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 2, b: 3})
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqA(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	if !setA.Eq(setA) {
		t.Fail()
	}
}

func Test_Comparable_EqB(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	if !setA.Eq(setA) {
		t.Fail()
	}
}

func Test_Comparable_EqC(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqD(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB := setA.Clone()
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqE(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqF(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setA.Add(testType_t{a: 16, b: 21})
	setB.Add(testType_t{a: 16, b: 21})
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqG(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 7, b: 5})
	setA.Add(testType_t{a: 7, b: 5})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqH(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB := setA.Clone()
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqI(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 16, b: 21})
	if setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqJ(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 7, b: 5})
	setB.Add(testType_t{a: 16, b: 21})
	setB.Add(testType_t{a: 7, b: 5})
	if setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointA(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	if !setA.IsDisjoint(setA) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointB(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointC(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointD(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 16, b: 21})
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointE(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 16, b: 21})
	setB.Add(testType_t{a: 7, b: 5})
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointF(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 16, b: 21})
	setA.Add(testType_t{a: 7, b: 5})
	setB := setA.Clone()
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointG(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 16, b: 21})
	setA.Add(testType_t{a: 7, b: 5})
	setB := setA.Clone()
	setB.Add(testType_t{a: 2, b: 3})
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointH(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 16, b: 21})
	setA.Add(testType_t{a: 7, b: 5})
	setB := setA.Clone()
	setA.Add(testType_t{a: 2, b: 3})
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_RemoveNonExisting(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	otherSet := mySet.Clone()
	mySet.Remove(testType_t{a: 16, b: 21})
	if !mySet.Eq(otherSet) {
		t.Fail()
	}
}

func Test_Comparable_RemoveItself(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	mySet.RemoveSet(mySet)
	if !mySet.IsEmpty() {
		t.Fail()
	}
}

func Test_Comparable_AddItselfA(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_Comparable_AddItselfB(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_Comparable_Representative(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	r := mySet.Representative(testType_t{a: 23, b: 14})
	if r.a != 23 || r.b != 14 {
		t.Fail()
	}
	r = mySet.Representative(testType_t{a: 16, b: 21})
	if r.a != 23 || r.b != 14 {
		t.Fail()
	}
	r = mySet.Representative(testType_t{a: 6, b: 6})
	if r.a != 7 || r.b != 5 {
		t.Fail()
	}
}

func Test_Comparable_RetainItselfA(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_Comparable_RetainItselfB(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_Comparable_RetainA(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 2, b: 3})
	setB := setA.Clone()
	setA.Add(testType_t{a: 7, b: 5})
	setA.Retain(setB)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_RetainB(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setA.Add(testType_t{a: 2, b: 3})
	setB := setA.Clone()
	setA.Add(testType_t{a: 7, b: 5})
	setA.Retain(setB)
	if !setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_Comparable_CardA(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 4, b: 1})
	mySet.Add(testType_t{a: 7, b: 5})
	mySet.Add(testType_t{a: 2, b: 3})
	if mySet.Card() != 2 {
		t.Fail()
	}
}

func Test_Comparable_Break(t *testing.T) {
	mySet := comparable.NewSet[testType_t]()
	mySet.Add(testType_t{a: 16, b: 21})
	mySet.Add(testType_t{a: 2, b: 3})
	mySet.Add(testType_t{a: 7, b: 5})
	g := mySet.Generator()
	resultList := []testType_t{}
	for elem := range g {
		if elem.CompareTo(testType_t{a: 11, b: 12}) < 0 {
			break
		}
		resultList = append(resultList, elem)
	}
	if len(resultList) > 2 {
		t.Fail()
	}
}
