package testseries

import "github.com/rhcourses-go/data-analysis/intlists"

// AbsoluteFrequencies erwartet eine Liste mit den Werten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den absoluten Häufigkeiten für jeden Wert
// zwischen dem Minimum und dem Maximum der Messreihe.
func AbsoluteFrequencies(values []int) []int {
	/* Hinweis:
	   Verwenden Sie die Funktionen Min und Max aus dem Paket intlists.
	   Erzeugen Sie eine Liste mit der richtigen Länge und fügen Sie in
	   einer Schleife über values die absoluten Häufigkeiten hinzu.

	   Dabei muss in jedem Schleifendurchlauf der Wert in der Ziel-Liste
	   um 1 erhöht werden, der an der Stelle steht, die dem Wert in der
	   Messreihe entspricht.

	   Eine Liste mit der richtigen Länge können Sie mit make erzeugen:
	   freq := make([]int, 5) erzeugt eine Liste mit 5 Elementen.
	*/
	// tag::solution[]
	min, max := intlists.Min(values), intlists.Max(values)
	freq := make([]int, max-min+1)
	for _, v := range values {
		freq[v-min] += 1
	}
	return freq
	// end::solution[]
	// taskreturn: []int{}
}

// RelativeFrequencies erwartet eine Liste mit absoluten Häufigkeiten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den relativen Häufigkeiten.
func RelativeFrequencies(values []int) []float64 {
	freq := make([]float64, len(values))
	/* Hinweis:
	   Bestimmen Sie die Summe aller absoluten Häufigkeiten
	   und berechnen Sie dann die relativen Häufigkeiten,
	   indem Sie in einer Schleife über die Werte iterieren
	   und in jedem Schleifendurchlauf den Wert durch die Summe teilen.
	*/
	// tag::solution[]
	sum := float64(intlists.Sum(values))
	for i, r := range values {
		freq[i] = float64(r) / sum
	}
	// end::solution[]
	return freq
}
