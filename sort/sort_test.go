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

var testcase = [...]int{100, 5, 15, 22, 66, 34, 15, 1, 1, 0}

func TestInsertionSort(t *testing.T) {
	x := testcase
	t.Log("before:", x)
	InsertionSort(x[0:])
	t.Log("after: ", x)
	if !isSorted(x[0:]) {
		t.Errorf("sort %v", testcase)
		t.Errorf(" got %v", x)
	}
}

func TestBubbleSort(t *testing.T) {
	x := testcase
	t.Log("before:", x)
	BubbleSort(x[0:])
	t.Log("after: ", x)
	if !isSorted(x[0:]) {
		t.Errorf("sort %v", testcase)
		t.Errorf(" got %v", x)
	}
}
