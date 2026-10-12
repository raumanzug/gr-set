package gr

import (
	"testing"

	"github.com/raumanzug/gr-set/comparable"
)

func Test_SortA(t *testing.T) {
	testContext["r"] = 7
	mySet := comparable.NewSet[testContextType_t, testType_t](
		testContext,
		testType_t{a: "r", b: "r"},
		testType_t{a: "x", b: "y"},
	)
	if mySet.Card() != 2 {
		t.Fail()
	}
	testContext["r"] = 6
	mySet.Sort()
	if mySet.Card() != 1 {
		t.Fail()
	}
}
