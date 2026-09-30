package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/comparable"
)

func Test_Comparable_VarArgs_A(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t](5, 12, 5)
	setC := comparable.NewSet[testType_t](12, 5)

	setA.Add(12)
	setA.Add(5)

	if !setA.Eq(setA) {
		t.Fail()
	}
	if !setA.Eq(setB) {
		t.Fail()
	}
	if !setA.Eq(setC) {
		t.Fail()
	}
	if !setB.Eq(setA) {
		t.Fail()
	}
	if !setB.Eq(setB) {
		t.Fail()
	}
	if !setB.Eq(setC) {
		t.Fail()
	}
	if !setC.Eq(setA) {
		t.Fail()
	}
	if !setC.Eq(setB) {
		t.Fail()
	}
	if !setC.Eq(setC) {
		t.Fail()
	}
}
