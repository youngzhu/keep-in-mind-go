package unsettled_test

import (
	"fmt"
	"github.com/youngzhu/keep-in-mind-go/unsettled"
)

var original = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func distributeEqually(original []int, n int) {
	fmt.Printf("Distribute %d parts\n", n)

	parts := unsettled.DistributeEqually(original, n)

	for i := range parts {
		fmt.Printf("  #%d: %v\n", i+1, parts[i])
	}
}

func ExampleDistributeEqually_3() {

	distributeEqually(original, 3)

	// Output:
	// Distribute 3 parts
	//   #1: [1 2 3]
	//   #2: [4 5 6]
	//   #3: [7 8 9 10]
}

func ExampleDistributeEqually_4() {

	distributeEqually(original, 4)

	// Output:
	// Distribute 4 parts
	//   #1: [1 2]
	//   #2: [3 4 5]
	//   #3: [6 7]
	//   #4: [8 9 10]
}

func ExampleDistributeEqually_5() {

	distributeEqually(original, 5)

	// Output:
	// Distribute 5 parts
	//   #1: [1 2]
	//   #2: [3 4]
	//   #3: [5 6]
	//   #4: [7 8]
	//   #5: [9 10]
}
