package tabledriven

import (
	"strconv"
	"testing"
)

func TestDaysOfMonth(t *testing.T) {
	testcases := []struct {
		month, days int
	}{
		{1, 31},
		{2, 28},
		{3, 31},
		{4, 30},
		{5, 31},
		{6, 30},
		{7, 31},
		{8, 31},
		{9, 30},
		{10, 31},
		{11, 30},
		{12, 31},
	}

	for i, tc := range testcases {
		no := strconv.Itoa(i + 1)
		t.Run("#"+no, func(t *testing.T) {
			got := DaysOfMonth(tc.month)
			if got != tc.days {
				t.Fatalf("got: %d, want: %d\n", got, tc.days)
			}
		})
	}
}
