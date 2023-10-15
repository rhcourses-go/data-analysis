package testseries

import "github.com/rhcourses-go/data-analysis/intlists"

// AbsoluteFrequencies erwartet eine Liste mit den Werten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den absoluten Häufigkeiten für jeden Wert
// zwischen dem Minimum und dem Maximum der Messreihe.
func AbsoluteFrequencies(values []int) []int {
	min, max := intlists.Min(values), intlists.Max(values)
	freq := make([]int, max-min+1)
	for _, v := range values {
		freq[v-min] += 1
	}
	return freq
}

// ValueCount erwartet eine Liste mit absoluten Häufigkeiten einer ganzzahligen Messreihe.
// Die Funktion liefert die Anzahl der Messwerte.
func ValueCount(values []int) int {
	totalValueCount := 0
	for _, r := range values {
		totalValueCount += r
	}
	return totalValueCount
}

// RelativeFrequencies erwartet eine Liste mit absoluten Häufigkeiten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den relativen Häufigkeiten.
func RelativeFrequencies(values []int) []float64 {
	freq := make([]float64, len(values))
	for i, r := range values {
		freq[i] = float64(r) / float64(ValueCount(values))
	}
	return freq
}
