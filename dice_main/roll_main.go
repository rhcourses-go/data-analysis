package main

import (
	"fmt"

	"github.com/rhcourses-go/data-analysis/dice"
	"github.com/rhcourses-go/data-analysis/testseries"
)

// readUserInput fragt den Benutzer nach der Anzahl der Würfe und der Anzahl der Würfel.
// Die Funktion liefert beide Werte zurück.
func readUserInput() (int, int) {
	fmt.Println("Würfelwurf-Simulator")
	fmt.Println("====================")
	fmt.Println()

	var n int
	fmt.Print("Wie oft soll gewürfelt werden? Bitte Zahl eingeben: ")
	fmt.Scanln(&n)

	var d int
	fmt.Print("Wie viele Würfel sollen verwendet werden? Bitte Zahl eingeben: ")
	fmt.Scanln(&d)

	fmt.Println()

	return d, n
}

func printDiceStatistics(rollResults []int) {
	n := len(rollResults)
	fmt.Println("Statistik für", n, "Würfelwürfe:")
	testseries.PrintDistribution(rollResults)
	fmt.Println()

	fmt.Println("Histogramm für", n, "Würfelwürfe:")
	testseries.PrintHistogram(rollResults)
	fmt.Println()
}

func main() {
	d, n := readUserInput()
	results := dice.RollMany(d, n)
	printDiceStatistics(results)
}
