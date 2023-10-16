package dice

import "math/rand"

// RollSingleDieOnce simuliert einen Würfelwurf: Die Funktion liefert eine Zufallszahl zwischen 1 und 6.
func RollSingleDieOnce() int {
	/* Hinweis:
	   Verwenden Sie die Funktion Intn aus dem Package math/rand.
	*/
	// tag::solution[]
	return rand.Intn(6) + 1
	// end::solution[]
	// taskreturn: return 0
}

// RollMultipleDiceOnce simuliert das einmalige Würfeln von zwei Würfeln.
// Die Funktion erwartet die Anzahl der Würfel und liefert die Summe der Würfelaugen.
func RollMultipleDiceOnce(d int) int {
	sum := 0
	/* Hinweis:
	   Schreiben Sie eine Schleife, die d mal die Funktion RollSingleDieOnce
	   aufruft und die Ergebnisse aufsummiert.
	*/
	// tag::solution[]
	for i := 0; i < d; i++ {
		sum += RollSingleDieOnce()
	}
	// end::solution[]
	return sum
}

// RollMany simuliert das mehrfache Würfeln eines oder mehrerer Würfel:
// Die Funktion erwartet die Anzahl d der Würfel und die Anzahl n der Würfe.
// Die Funktion würfelt n mal und liefert eine Liste mit den Ergebnissen.
func RollMany(d, n int) []int {
	rollResults := make([]int, n)
	/* Hinweis:
	   Schreiben Sie eine Schleife, die n mal die Funktion RollMultipleDiceOnce
	   aufruft und die Ergebnisse in der Liste rollResults speichert.
	*/
	// tag::solution[]
	for i := 0; i < n; i++ {
		rollResults[i] = RollMultipleDiceOnce(d)
	}
	// end::solution[]
	return rollResults
}

// Anmerkung: Die Funktionen RollOnce und RollMany sind nicht deterministisch.
// Das heißt, sie liefern bei jedem Aufruf ein anderes Ergebnis.
//
// Diese Funktionen können daher nicht auf einfache Art getestet werden,
// da die Testergebnisse nicht vorhersehbar sind.
// Um Zufallsfunktionen zu testen, müsste man ihnen als Parameter einen
// Zufallsgenerator übergeben, den sie benutzen.
// Diesen könnte man dann im Test durch einen deterministischen Zufallsgenerator
// ersetzen, der immer die gleichen Zahlen liefert.
//
// Das würde hier aber den Rahmen sprengen,
// daher verzichten wir bei diesem Package auf Tests.
