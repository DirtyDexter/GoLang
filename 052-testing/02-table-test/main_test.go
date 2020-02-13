package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{4, 5, 6}, 15},
		test{[]int{7, 8, 9}, 24},
		test{[]int{10, 11, 12}, 33},
	}

	for _, v := range tests {
		s := sum(v.data...)
		if s != v.answer {
			t.Error("expected", v.answer, "got", s)
		}
	}
}
