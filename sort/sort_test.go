package sort

import (
	"testing"
)

// testdata
// array 保证每次排序前都是无序的
var testcase = [...]int{100, 5, 15, 22, 66, 34, 15, 1, 1, 0}

//var testcase = [...]int{1, 2, 3, 4}
//var testcase = [...]int{4, 3, 2, 1}

func isSorted(x []int) bool {
	for i := 1; i < len(x); i++ {
		if x[i] < x[i-1] {
			return false
		}
	}
	return true
}

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

func TestMergeSort(t *testing.T) {
	x := testcase
	MergeSort(x[0:])
	if !isSorted(x[0:]) {
		t.Errorf("sort %v", testcase)
		t.Errorf(" got %v", x)
	}
}
