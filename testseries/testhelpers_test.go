package testseries

import "fmt"

func printFloatList(list []float64) {
	for _, v := range list[:len(list)-1] {
		fmt.Printf("%.2f ", v)
	}
	fmt.Printf("%.2f\n", list[len(list)-1])
}
