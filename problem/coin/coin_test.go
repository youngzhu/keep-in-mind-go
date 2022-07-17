package coin

import (
	"strconv"
	"testing"
)

func TestCoin(t *testing.T) {
	t.Log(5 << 0)
	t.Log(5 << 1)
	t.Log(5 << 2)
}

func TestChanges_simple(t *testing.T) {
	Changes(1)
}

func TestChanges(t *testing.T) {
	testcases := []struct {
		n, result int
	}{
		{1, 1},
		{5, 2},
		{10, 4},
		{15, 6},
		{20, 9},
		{25, 13},
		{30, 18},
		{35, 24},
		{40, 31},
		{45, 39},
		{50, 50},
		{100, 292},
	}

	for i, tc := range testcases {
		no := strconv.Itoa(i + 1)
		t.Run("#"+no, func(t *testing.T) {
			got := Changes(tc.n)
			want := tc.result
			if got != want {
				t.Fatalf("got: %d, want: %d\n", got, want)
			}
		})
	}
}
