package main

import "testing"

func TestSum(t *testing.T) {
	s := sum(2, 3, 4, 5)
	if s != 14 {
		t.Error("expected", 14, "got", s)
	}
}
