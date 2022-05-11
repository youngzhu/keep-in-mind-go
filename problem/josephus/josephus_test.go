package josephus

import (
	"reflect"
	"strconv"
	"testing"
)

var testcases = []struct {
	m, n   int
	result []int
}{
	{2, 7, []int{1, 3, 5, 0, 4, 2, 6}},
}

func TestSolve(t *testing.T) {
	for i, tc := range testcases {
		no := strconv.Itoa(i + 1)
		t.Run("#"+no, func(t *testing.T) {
			got := Solve(tc.m, tc.n)
			want := tc.result
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %d, want: %d\n", got, want)
			}
		})
	}
}
