package testseries

import "fmt"

func ExampleAbsoluteFrequencies() {
	rawData1 := []int{1, 3, 4, 1, 2, 1, 5, 7, 5, 4, 2}
	// Value: 1 | 2 | 3 | 4 | 5 | 6 | 7
	// Count: 3 | 2 | 1 | 2 | 2 | 0 | 1
	fmt.Println(AbsoluteFrequencies(rawData1))

	rawData2 := []int{15, 12, 10, 12, 10, 13}
	// Value: 10 | 11 | 12 | 13 | 14 | 15
	// Count:  2 |  0 |  2 |  1 |  0 |  1
	fmt.Println(AbsoluteFrequencies(rawData2))

	// Output:
	// [3 2 1 2 2 0 1]
	// [2 0 2 1 0 1]
}

func ExampleRelativeFrequencies() {
	absFreqs1 := []int{3, 2, 1, 2, 2, 0, 1}
	// Value:    1 |    2 |    3 |    4 |    5 |    6 |    7
	// Count:    3 |    2 |    1 |    2 |    2 |    0 |    1
	// Relat: 0.27 | 0.18 | 0.09 | 0.18 | 0.18 | 0.00 | 0.09
	printFloatList(RelativeFrequencies(absFreqs1))

	absFreqs2 := []int{2, 0, 2, 1, 0, 1}
	// Value:    10 |    11 |    12 |    13 |    14 |    15
	// Count:     2 |     0 |     2 |     1 |     0 |     1
	// Relat: 0.33 | 0.00 | 0.33 | 0.17 | 0.00 | 0.17
	printFloatList(RelativeFrequencies(absFreqs2))

	// Output:
	// 0.27 0.18 0.09 0.18 0.18 0.00 0.09
	// 0.33 0.00 0.33 0.17 0.00 0.17
}
