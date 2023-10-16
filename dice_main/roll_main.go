package main

import (
	"fmt"

	"github.com/rhcourses-go/data-analysis/dice"
	"github.com/rhcourses-go/data-analysis/testseries"
)

// readUserInput fragt den Benutzer nach der Anzahl der Würfe und der Anzahl der Würfel.
// Die Funktion liefert beide Werte zurück.
func readUserInput() (int, int) {
	var d, n int
	/* Hinweis:
	   Verwenden Sie die Funktionen Println und Scanln aus dem Package fmt,
	   um Text auf der Konsole auszugeben und Eingaben vom Benutzer einzulesen.

	   Speichern Sie die Eingaben in den Variablen n und d.
	*/
	// tag::solution[]
	fmt.Println("Würfelwurf-Simulator")
	fmt.Println("====================")
	fmt.Println()

	fmt.Print("Wie oft soll gewürfelt werden? Bitte Zahl eingeben: ")
	fmt.Scanln(&n)

	fmt.Print("Wie viele Würfel sollen verwendet werden? Bitte Zahl eingeben: ")
	fmt.Scanln(&d)

	fmt.Println()
	// end::solution[]
	return d, n
}

// printDiceStatistics berechnet die Statistik für die Würfelwürfe und gibt sie aus.
func printDiceStatistics(rollResults []int) {
	/* Hinweis:
	   Verwenden Sie die Funktionen PrintDistribution und PrintHistogram
	   aus dem Package testseries, um die Statistik auszugeben.
	*/
	// tag::solution[]
	n := len(rollResults)
	fmt.Println("Statistik für", n, "Würfelwürfe:")
	testseries.PrintDistribution(rollResults)
	fmt.Println()

	fmt.Println("Histogramm für", n, "Würfelwürfe:")
	testseries.PrintHistogram(rollResults)
	fmt.Println()
	// end::solution[]
}

// main kombiniert die anderen Funktionen zu einem Programm.
func main() {
	// tag::solution[]
	d, n := readUserInput()
	results := dice.RollMany(d, n)
	printDiceStatistics(results)
	// end::solution[]
}
