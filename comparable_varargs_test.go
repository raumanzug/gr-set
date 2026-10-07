package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/comparable"
)

func Test_Comparable_VarArgs_A(t *testing.T) {
	setA := comparable.NewSet[testType_t]()
	setB := comparable.NewSet[testType_t](
		testType_t{a: 3, b: 2},
		testType_t{a: 6, b: 6},
		testType_t{a: 1, b: 4})
	setC := comparable.NewSet[testType_t](
		testType_t{a: 6, b: 6},
		testType_t{a: 3, b: 2},
	)

	setA.Add(testType_t{a: 6, b: 6})
	setA.Add(testType_t{a: 3, b: 2})

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
