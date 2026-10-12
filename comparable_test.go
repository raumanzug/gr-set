package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/comparable"
)

type testContextType_t map[string]uint

type testType_t struct {
	a string
	b string
}

func (recv testType_t) val(ctxt testContextType_t) uint {
	ai := ctxt[recv.a]
	bi := ctxt[recv.b]
	return ai + bi
}

var testContext = testContextType_t{
	"a": 1,
	"b": 4,
	"m": 11,
	"n": 12,
	"p": 16,
	"q": 21,
	"r": 6,
	"s": 14,
	"t": 23,
	"u": 2,
	"v": 3,
	"x": 5,
	"y": 7,
	"z": 37,
}

func (recv testType_t) CompareTo(ctxt testContextType_t, other testType_t) int {
	sum_left := recv.val(ctxt)
	sum_right := other.val(ctxt)
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
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	if !mySet.IsEmpty() {
		t.Fatalf("freshly initialized comparable set should be empty.  is indeed not.")
	}
}

func Test_Comparable_IsNonEmpty(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "y", b: "x"})
	if mySet.IsEmpty() {
		t.Fatalf("set after adding some elements should not be empty.  is indeed empty.")
	}
}

func Test_Comparable_IsSubeqA(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsSubeqB(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB := setA.Clone()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsSubeqC(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := setA.Clone()
	setB.Add(testType_t{a: "y", b: "x"})
	if setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_Comparable_AddIdempotenceA(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_AddIdempotenceB(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_AddCommutativity(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "u", b: "v"})
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "u", b: "v"})
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqA(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	if !setA.Eq(setA) {
		t.Fail()
	}
}

func Test_Comparable_EqB(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	if !setA.Eq(setA) {
		t.Fail()
	}
}

func Test_Comparable_EqC(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqD(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB := setA.Clone()
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqE(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqF(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setA.Add(testType_t{a: "p", b: "q"})
	setB.Add(testType_t{a: "p", b: "q"})
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqG(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "y", b: "x"})
	setA.Add(testType_t{a: "y", b: "x"})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqH(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB := setA.Clone()
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqI(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "p", b: "q"})
	if setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_EqJ(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "y", b: "x"})
	setB.Add(testType_t{a: "p", b: "q"})
	setB.Add(testType_t{a: "y", b: "x"})
	if setA.Eq(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointA(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	if !setA.IsDisjoint(setA) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointB(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointC(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointD(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "p", b: "q"})
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointE(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "p", b: "q"})
	setB.Add(testType_t{a: "y", b: "x"})
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointF(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "p", b: "q"})
	setA.Add(testType_t{a: "y", b: "x"})
	setB := setA.Clone()
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointG(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "p", b: "q"})
	setA.Add(testType_t{a: "y", b: "x"})
	setB := setA.Clone()
	setB.Add(testType_t{a: "u", b: "v"})
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_IsDisjointH(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "p", b: "q"})
	setA.Add(testType_t{a: "y", b: "x"})
	setB := setA.Clone()
	setA.Add(testType_t{a: "u", b: "v"})
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_Comparable_RemoveNonExisting(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	otherSet := mySet.Clone()
	mySet.Remove(testType_t{a: "p", b: "q"})
	if !mySet.Eq(otherSet) {
		t.Fail()
	}
}

func Test_Comparable_RemoveItself(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	mySet.RemoveSet(mySet)
	if !mySet.IsEmpty() {
		t.Fail()
	}
}

func Test_Comparable_AddItselfA(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_Comparable_AddItselfB(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_Comparable_Representative(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	r := mySet.Representative(testType_t{a: "t", b: "s"})
	if r.a != "t" || r.b != "s" {
		t.Fail()
	}
	r = mySet.Representative(testType_t{a: "p", b: "q"})
	if r.a != "t" || r.b != "s" {
		t.Fail()
	}
	r = mySet.Representative(testType_t{a: "r", b: "r"})
	if r.a != "y" || r.b != "x" {
		t.Fail()
	}
}

func Test_Comparable_RetainItselfA(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_Comparable_RetainItselfB(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_Comparable_RetainA(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "u", b: "v"})
	setB := setA.Clone()
	setA.Add(testType_t{a: "y", b: "x"})
	setA.Retain(setB)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_Comparable_RetainB(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setA.Add(testType_t{a: "u", b: "v"})
	setB := setA.Clone()
	setA.Add(testType_t{a: "y", b: "x"})
	setA.Retain(setB)
	if !setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_Comparable_CardA(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "b", b: "a"})
	mySet.Add(testType_t{a: "y", b: "x"})
	mySet.Add(testType_t{a: "u", b: "v"})
	if mySet.Card() != 2 {
		t.Fail()
	}
}

func Test_Comparable_Break(t *testing.T) {
	mySet := comparable.NewSet[testContextType_t, testType_t](testContext)
	mySet.Add(testType_t{a: "p", b: "q"})
	mySet.Add(testType_t{a: "u", b: "v"})
	mySet.Add(testType_t{a: "y", b: "x"})
	g := mySet.Generator()
	resultList := []testType_t{}
	for elem := range g {
		if elem.CompareTo(testContext, testType_t{a: "m", b: "n"}) < 0 {
			break
		}
		resultList = append(resultList, elem)
	}
	if len(resultList) > 2 {
		t.Fail()
	}
}
