package testseries

import (
	"fmt"

	"github.com/rhcourses-go/data-analysis/intlists"
)

// PrintDistribution erwartet eine Liste mit Ergebnissen einer Integer-Messreihe.
// Die Funktion gibt die absolute und relative Häufigkeit sowie
// den Wert der empirischen Verteilungsfunktion für jede Zahl aus.
func PrintDistribution(values []int) {
	/* Hinweis:
	   Verwenden Sie die Funktionen aus den vorherigen Aufgaben,
	   um die absoluten und relativen Häufigkeiten sowie die
	   empirische Verteilungsfunktion zu berechnen.

	   Geben Sie die Werte in einer for-Schleife aus,
	   die über die Werte der Messreihe iteriert.
	   Beachten Sie dabei die Abstände zwischen den Spalten.
	*/
	// tag::solution[]
	valrange := intlists.ValueRange(values)
	absolute := AbsoluteFrequencies(values)
	relative := RelativeFrequencies(absolute)
	emp := Distribution(values)

	fmt.Println("Wert   Abs.   Rel.     Vert.")

	for i, v := range valrange {
		fmt.Printf("%d       %d     %.2f     %.2f\n", v, absolute[i], relative[i], emp[i])
	}
	// end::solution[]
}

// PrintHistogram erwartet eine Liste mit Ergebnissen einer Integer-Messreihe.
// Die Funktion gibt ein Histogramm aus.
func PrintHistogram(values []int) {
	/* Hinweis:
	   Verwenden Sie die Funktionen aus den vorherigen Aufgaben,
	   um den Wertebereich und die absoluten Häufigkeiten zu berechnen.

	   Mit diesen beiden Listen können Sie dann in einer for-Schleife
	   über den Wertebereich iterieren und in jeder Iteration
	   die entsprechende Anzahl an Sternchen ausgeben.
	*/
	// tag::solution[]
	valrange := intlists.ValueRange(values)
	absolute := AbsoluteFrequencies(values)

	for i, v := range valrange {
		fmt.Printf("%d: |", v)
		for j := 0; j < absolute[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// end::solution[]
}
