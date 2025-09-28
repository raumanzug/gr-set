package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/simple"
)

func Test_IsEmpty(t *testing.T) {
	mySet := simple.NewSet[uint]()
	if !mySet.IsEmpty() {
		t.Fatalf("freshly initialized simple set should be empty.  is indeed not.")
	}
}

func Test_IsNonEmpty(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(12)
	if mySet.IsEmpty() {
		t.Fatalf("set after adding some elements should not be empty.  is indeed empty.")
	}
}

func Test_IsSubeqA(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_IsSubeqB(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(12)
	setB := setA.Clone()
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_IsSubeqC(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := setA.Clone()
	setB.Add(12)
	if setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_AddIdempotenceA(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setB.Add(12)
	setB.Add(12)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_AddIdempotenceB(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setA.Add(12)
	setB.Add(12)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_AddCommutativity(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(5)
	setA.Add(12)
	setB.Add(12)
	setB.Add(5)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_EqA(t *testing.T) {
	setA := simple.NewSet[uint]()
	if !setA.Eq(setA) {
		t.Fail()
	}
}

func Test_EqB(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(12)
	if !setA.Eq(setA) {
		t.Fail()
	}
}

func Test_EqC(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqD(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(12)
	setB := setA.Clone()
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqE(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setB.Add(12)
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqF(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setA.Add(37)
	setB.Add(37)
	setB.Add(12)
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqG(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setB.Add(12)
	setA.Add(12)
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqH(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(12)
	setB := setA.Clone()
	setB.Add(12)
	if !setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqI(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setB.Add(37)
	if setA.Eq(setB) {
		t.Fail()
	}
}

func Test_EqJ(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(12)
	setB.Add(37)
	setB.Add(12)
	if setA.Eq(setB) {
		t.Fail()
	}
}

func Test_IsDisjointA(t *testing.T) {
	setA := simple.NewSet[uint]()
	if !setA.IsDisjoint(setA) {
		t.Fail()
	}
}

func Test_IsDisjointB(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_IsDisjointC(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setB.Add(12)
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_IsDisjointD(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(37)
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_IsDisjointE(t *testing.T) {
	setA := simple.NewSet[uint]()
	setB := simple.NewSet[uint]()
	setA.Add(37)
	setB.Add(12)
	if !setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_IsDisjointF(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(37)
	setA.Add(12)
	setB := setA.Clone()
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_IsDisjointG(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(37)
	setA.Add(12)
	setB := setA.Clone()
	setB.Add(5)
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_IsDisjointH(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(37)
	setA.Add(12)
	setB := setA.Clone()
	setA.Add(5)
	if setA.IsDisjoint(setB) {
		t.Fail()
	}
}

func Test_RemoveItself(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySet.RemoveSet(mySet)
	if !mySet.IsEmpty() {
		t.Fail()
	}
}

func Test_AddItselfA(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_AddItselfB(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.AddSet(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_RetainItselfA(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySet.Subeq(mySetClone) {
		t.Fail()
	}
}

func Test_RetainItselfB(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySetClone := mySet
	mySet.Retain(mySet)
	if !mySetClone.Subeq(mySet) {
		t.Fail()
	}
}

func Test_RetainA(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(5)
	setB := setA.Clone()
	setA.Add(12)
	setA.Retain(setB)
	if !setA.Subeq(setB) {
		t.Fail()
	}
}

func Test_RetainB(t *testing.T) {
	setA := simple.NewSet[uint]()
	setA.Add(5)
	setB := setA.Clone()
	setA.Add(12)
	setA.Retain(setB)
	if !setB.Subeq(setA) {
		t.Fail()
	}
}

func Test_CardA(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(5)
	mySet.Add(12)
	mySet.Add(5)
	if mySet.Card() != 2 {
		t.Fail()
	}
}

func Test_Break(t *testing.T) {
	mySet := simple.NewSet[uint]()
	mySet.Add(37)
	mySet.Add(5)
	mySet.Add(12)
	g := mySet.Generator()
	resultList := []uint{}
	for elem := range g {
		if elem < 23 {
			break
		}
		resultList = append(resultList, elem)
	}
	if len(resultList) > 2 {
		t.Fail()
	}
}
