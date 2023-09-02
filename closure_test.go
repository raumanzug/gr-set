package set

import (
	"github.com/raumanzug/gr-generator"
	"testing"
)

var (
	connections = map[string][]string{
		"Dresden":  {"Chemnitz", "Leipzig"},
		"Chemnitz": {"Dresden", "Leipzig", "Hof"},
		"Hof":      {"Hof"},
		"Leipzig":  {"Halle"},
		"Halle":    {"Magdeburg"},
	}
)

func op(elem string) generator.Generator[generator.LoopDirective, string] {
	return generator.Array2Generator(connections[elem])
}

func Test_EmptyClosure(t *testing.T) {
	base := NewSimpleSet[string]()

	Closure[string](base, op)

	if !base.IsEmpty() {
		t.Fail()
	}
}

func Test_SingleClosureA(t *testing.T) {
	base := NewSimpleSet[string]()
	base.Add("Hof")
	cmp := base.Clone()

	Closure[string](base, op)

	if !base.Subeq(cmp) || !cmp.Subeq(base) {
		t.Fail()
	}
}

func Test_SingleClosureB(t *testing.T) {
	base := NewSimpleSet[string]()
	base.Add("Magdeburg")
	cmp := base.Clone()

	Closure[string](base, op)

	if !base.Subeq(cmp) || !cmp.Subeq(base) {
		t.Fail()
	}
}

func Test_singleClosureC(t *testing.T) {
	base := NewSimpleSet[string]()
	base.Add("Leipzig")

	Closure[string](base, op)

	if !base.Contains("Leipzig") ||
		!base.Contains("Halle") ||
		!base.Contains("Magdeburg") ||
		base.Card() != 3 {
		t.Fail()
	}
}

func Test_singleClosureD(t *testing.T) {
	base := NewSimpleSet[string]()
	base.Add("Dresden")

	Closure[string](base, op)

	if !base.Contains("Dresden") ||
		!base.Contains("Chemnitz") ||
		!base.Contains("Leipzig") ||
		!base.Contains("Halle") ||
		!base.Contains("Magdeburg") ||
		!base.Contains("Hof") ||
		base.Card() != 6 {
		t.Fail()
	}
}

func Test_doubleClosureA(t *testing.T) {
	base := NewSimpleSet[string]()
	base.Add("Dresden")
	base.Add("Hof")

	Closure[string](base, op)

	if !base.Contains("Dresden") ||
		!base.Contains("Chemnitz") ||
		!base.Contains("Leipzig") ||
		!base.Contains("Halle") ||
		!base.Contains("Magdeburg") ||
		!base.Contains("Hof") ||
		base.Card() != 6 {
		t.Fail()
	}
}

func Test_doubleClosureB(t *testing.T) {
	base := NewSimpleSet[string]()
	base.Add("Leipzig")
	base.Add("Hof")

	Closure[string](base, op)

	if !base.Contains("Leipzig") ||
		!base.Contains("Halle") ||
		!base.Contains("Magdeburg") ||
		!base.Contains("Hof") ||
		base.Card() != 4 {
		t.Fail()
	}
}
