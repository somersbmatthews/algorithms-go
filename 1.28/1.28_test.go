package main

import "testing"

var data = []int{1, 3, 4}

func TestSumArray(t *testing.T) {
	got := SumArray(data)
	want := 8
	if got != want {
		t.Errorf("SumArray function is broken, got %d, want %d", got, want)
	}
}
