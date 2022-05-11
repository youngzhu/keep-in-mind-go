package josephus_test

import (
	"fmt"
	"github.com/youngzhu/keep-in-mind-go/problem/josephus"
)

func ExampleSolve() {
	var result []int

	result = josephus.Solve(2, 7)
	fmt.Println(result)

	// Output:
	// [1 3 5 0 4 2 6]
}
