package intlists

import "fmt"

func ExampleMin() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Min(list1))

	list2 := []int{3, 4, 5}
	fmt.Println(Min(list2))

	list3 := []int{5, 3, 4}
	fmt.Println(Min(list3))

	// Output:
	// 1
	// 3
	// 3
}

func ExampleMax() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Max(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(Max(list2))

	list3 := []int{5, 3, 4}
	fmt.Println(Max(list3))

	// Output:
	// 5
	// 6
	// 5
}

func ExampleValueRange() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(ValueRange(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(ValueRange(list2))

	list3 := []int{7, 3, 2}
	fmt.Println(ValueRange(list3))

	// Output:
	// [1 2 3 4 5]
	// [3 4 5 6]
	// [2 3 4 5 6 7]
}

func ExampleSum() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(Sum(list2))

	list3 := []int{}
	fmt.Println(Sum(list3))

	// Output:
	// 15
	// 13
	// 0
}

func ExampleProduct() {
	list1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Product(list1))

	list2 := []int{3, 4, 6}
	fmt.Println(Product(list2))

	list3 := []int{}
	fmt.Println(Product(list3))

	// Output:
	// 120
	// 72
	// 1
}
