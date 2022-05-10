package sort

import (
	"testing"
)

func isSorted(x []int) bool {
	for i := 1; i < len(x); i++ {
		if x[i] < x[i-1] {
			return false
		}
	}
	return true
}

var testcase = []int{100, 5, 15, 22, 66, 34, 15, 1, 1, 0}

func TestInsertionSort(t *testing.T) {
	t.Log("before:", testcase)
	InsertionSort(testcase)
	t.Log("after: ", testcase)
	if !isSorted(testcase) {
		t.Fatal("sort fail!")
	}
}
