package main

import "testing"

func TestFirst(t * testing.T) {
	a := []int64{1,2}
	r := SumOfBits(a)
	if r != 4 {
		t.Error("expected", 4, "got", r)
	}
}