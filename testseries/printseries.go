package testseries

import (
	"fmt"

	"github.com/rhcourses-go/data-analysis/intlists"
)

// PrintDistribution erwartet eine Liste mit Ergebnissen einer Integer-Messreihe.
// Die Funktion gibt die absolute und relative Häufigkeit sowie
// den Wert der empirischen Verteilungsfunktion für jede Zahl aus.
func PrintDistribution(values []int) {
	valrange := intlists.ValueRange(values)
	absolute := AbsoluteFrequencies(values)
	relative := RelativeFrequencies(absolute)
	emp := Distribution(values)

	fmt.Println("Wert   Abs.   Rel.     Vert.")

	for i, v := range valrange {
		fmt.Printf("%d       %d     %.2f     %.2f\n", v, absolute[i], relative[i], emp[i])
	}
}

// PrintHistogram erwartet eine Liste mit Ergebnissen einer Integer-Messreihe.
// Die Funktion gibt ein Histogramm aus.
func PrintHistogram(values []int) {
	valrange := intlists.ValueRange(values)
	absolute := AbsoluteFrequencies(values)

	for i, v := range valrange {
		fmt.Printf("%d: |", v)
		for j := 0; j < absolute[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
