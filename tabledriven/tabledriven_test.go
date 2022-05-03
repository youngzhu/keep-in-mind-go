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

func TestGetGrade(t *testing.T) {
	testcases := []struct {
		score float32
		grade rune
	}{
		{100, 'A'},
		{99.5, 'A'},
		{90, 'A'},
		{89.5, 'B'},
		{80, 'B'},
		{75, 'B'},
		{74.5, 'C'},
		{69, 'C'},
		{65, 'C'},
		{64.5, 'D'},
		{50, 'D'},
		{49.5, 'F'},
		{49, 'F'},
	}

	for i, tc := range testcases {
		no := strconv.Itoa(i + 1)
		t.Run("#"+no, func(t *testing.T) {
			got := GetGrade(tc.score)
			want := tc.grade
			if got != want {
				t.Fatalf("got: %d, want: %d\n", got, want)
			}
		})
	}
}
