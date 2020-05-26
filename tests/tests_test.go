package main

import (
	"fmt"
	"os"
	"testing"
)

var shared string

func TestMain(m *testing.M) {
	shared = "test main 1"
	fmt.Println("Test Main 1")
	os.Exit(m.Run())
}

func TestA(t *testing.T) {
	t.Logf("Test A, shared = %v", shared)
}

func TestUsualSum(t *testing.T) {
	testCases := []struct {
		x, y, want int
	}{
		{0, 1, 1},
		{0, 2, 2},
		{2, 0, 2},
	}

	for _, tc := range testCases {
		if got := sum(tc.x, tc.y); got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestSubTestSum(t *testing.T) {

}

func BenchmarkSum(b *testing.B) {
	x, y := 3, 7
	for i := 0; i < b.N; i++ {
		_ = sum(x, y)
	}
}
