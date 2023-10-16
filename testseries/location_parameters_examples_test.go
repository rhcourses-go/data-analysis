package testseries

import "fmt"

func ExampleAverage() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Average(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(Average(list2))

	list3 := []int{}
	fmt.Println(Average(list3))

	list4 := []int{1, 1, 1}
	fmt.Println(Average(list4))

	// Output:
	// 3
	// 4.333333333333333
	// 0
	// 1
}

func ExampleMedian() {
	list1 := []int{10, 20, 30, 40, 50}
	fmt.Println(Median(list1))

	list2 := []int{50, 40, 30, 20}
	fmt.Println(Median(list2))

	list3 := []int{1, 17, 2, 42, 38, 3, 5}
	fmt.Println(Median(list3))

	list4 := []int{}
	fmt.Println(Median(list4))

	// Output:
	// 30
	// 35
	// 5
	// 0
}

func ExampleMode() {
	list1 := []int{1, 2, 3, 4, 5} // Alle gleich häufig.
	fmt.Println(Mode(list1))

	list2 := []int{2, 3, 4, 5, 1} // Alle gleich häufig.
	fmt.Println(Mode(list2))

	list3 := []int{1, 1, 2, 3, 4} // 1 am häufigsten.
	fmt.Println(Mode(list3))

	list4 := []int{1, 2, 2, 3, 4} // 2 am häufigsten.
	fmt.Println(Mode(list4))

	list5 := []int{4, 2, 3, 4, 5, 2, 3, 7, 5, 3, 1, 1, 2, 5, 7, 5, 38} // 5 am häufigsten.
	fmt.Println(Mode(list5))

	list6 := []int{} // leer
	fmt.Println(Mode(list6))

	// Output:
	// 1
	// 1
	// 1
	// 2
	// 5
	// 0
}

func ExampleGeometricMean() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(GeometricMean(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(GeometricMean(list2))

	list3 := []int{}
	fmt.Println(GeometricMean(list3))

	list4 := []int{1, 1, 1}
	fmt.Println(GeometricMean(list4))

	list5 := []int{3, 300}
	fmt.Println(GeometricMean(list5))

	// Output:
	// 2.605171084697352
	// 4.160167646103807
	// 0
	// 1
	// 30
}

func ExampleHarmonicMean() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(HarmonicMean(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(HarmonicMean(list2))

	list3 := []int{}
	fmt.Println(HarmonicMean(list3))

	list4 := []int{1, 1, 1}
	fmt.Println(HarmonicMean(list4))

	list5 := []int{5, 20}
	fmt.Println(HarmonicMean(list5))

	// Output:
	// 2.18978102189781
	// 4.000000000000001
	// 0
	// 1
	// 8
}
