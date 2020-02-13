package mypkg

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	s := Sum(2, 3, 4, 5)
	if s != 14 {
		t.Error("expected", 14, "got", s)
	}
}

func ExampleSum() {
	fmt.Println(Sum(2, 3, 4))
	// Output:
	// 9
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(10, 3445, 345, 12)
	}
}
