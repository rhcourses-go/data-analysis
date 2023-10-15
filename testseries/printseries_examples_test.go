package testseries

import "fmt"

func ExamplePrintDistribution() {
	rawData1 := []int{1, 3, 4, 1, 2, 1, 5, 7, 5, 4, 2}
	PrintDistribution(rawData1)
	fmt.Println()

	rawData2 := []int{15, 12, 10, 12, 10, 13}
	PrintDistribution(rawData2)
	fmt.Println()

	// Output:
	// Wert   Abs.   Rel.     Vert.
	// 1       3     0.27     0.27
	// 2       2     0.18     0.45
	// 3       1     0.09     0.55
	// 4       2     0.18     0.73
	// 5       2     0.18     0.91
	// 6       0     0.00     0.91
	// 7       1     0.09     1.00
	//
	// Wert   Abs.   Rel.     Vert.
	// 10       2     0.33     0.33
	// 11       0     0.00     0.33
	// 12       2     0.33     0.67
	// 13       1     0.17     0.83
	// 14       0     0.00     0.83
	// 15       1     0.17     1.00
}

func ExamplePrintHistogram() {
	rawData1 := []int{1, 3, 4, 1, 2, 1, 4, 5, 7, 5, 4, 2, 4}
	PrintHistogram(rawData1)

	fmt.Println("---")

	rawData2 := []int{15, 12, 10, 12, 10, 13, 12, 14}
	PrintHistogram(rawData2)

	// Output:
	// 1: |***
	// 2: |**
	// 3: |*
	// 4: |****
	// 5: |**
	// 6: |
	// 7: |*
	// ---
	// 10: |**
	// 11: |
	// 12: |***
	// 13: |*
	// 14: |*
	// 15: |*
}
