package main

import "testing"

func TestSequentialSearch(t *testing.T) {
	boolGot, indexGot := SequentialSearch(data, search)
	boolWant, indexWant := true, 1
	if boolGot != boolWant {
		t.Errorf("Sequential search is broken. got: %t but want: %t", boolGot, boolWant)
	}
	if indexGot != indexWant {
		t.Errorf("Sequential search is broken. got: %d want: %d", indexGot, indexWant)
	}
}
