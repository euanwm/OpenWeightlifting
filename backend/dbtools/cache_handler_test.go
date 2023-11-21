package dbtools

import "testing"

func TestCompareStructs(t *testing.T) {
	type TestOne struct {
		One int
		Two string
	}

	testOne := TestOne{One: 1, Two: "two"}
	testTwo := TestOne{One: 1, Two: "two"}

	if testOne != testTwo {
		t.Errorf("Structs are not equal")
	}
}
