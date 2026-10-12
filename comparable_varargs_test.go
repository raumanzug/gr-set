package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/comparable"
)

func Test_Comparable_VarArgs_A(t *testing.T) {
	setA := comparable.NewSet[testContextType_t, testType_t](testContext)
	setB := comparable.NewSet[testContextType_t, testType_t](
		testContext,
		testType_t{a: "v", b: "u"},
		testType_t{a: "r", b: "r"},
		testType_t{a: "a", b: "b"})
	setC := comparable.NewSet[testContextType_t, testType_t](
		testContext,
		testType_t{a: "r", b: "r"},
		testType_t{a: "v", b: "u"},
	)

	setA.Add(testType_t{a: "r", b: "r"})
	setA.Add(testType_t{a: "v", b: "u"})

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
