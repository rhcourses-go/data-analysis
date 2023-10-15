package testseries

import (
	"math"
	"slices"

	"github.com/rhcourses-go/data-analysis/intlists"
)

// Average erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Durchschnittswert.
// Ist die Liste leer, wird 0.0 zurückgegeben.
func Average(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}
	return float64(intlists.Sum(values)) / float64(len(values))
}

// Median erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Median.
// Ist die Liste leer, wird 0 zurückgegeben.
func Median(values []int) int {
	if len(values) == 0 {
		return 0
	}

	sorted := make([]int, len(values))
	copy(sorted, values)
	slices.Sort(sorted)

	if len(sorted)%2 == 0 {
		return (sorted[len(sorted)/2-1] + sorted[len(sorted)/2]) / 2
	}
	return sorted[len(sorted)/2]
}

// Mode erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Wert, der am häufigsten in der Liste vorkommt.
// Falls mehrere Werte am häufigsten vorkommen, wird der kleinste dieser Werte
// zurückgegeben. Ist die Liste leer, wird 0 zurückgegeben.
func Mode(values []int) int {
	if len(values) == 0 {
		return 0
	}

	// Bestimme den kleinsten Wert, der in den Daten vorkommt.
	minvalue := intlists.Min(values)

	// Bestimme, wie oft der häufigste Wert vorkommt.
	// Bestimme dazu die Liste der absoluten Häufigkeiten und deren Maximum.
	freq := AbsoluteFrequencies(values)
	max := intlists.Max(freq)

	// Suche die Position dieses Wertes in der Liste der absoluten Häufigkeiten.
	// Aus dieser Position kann der häufigste Wert bestimmt werden.
	for i, f := range freq {
		if f == max {
			return i + minvalue
		}
	}

	// Wenn diese Zeile erreicht wird, ist ein Fehler aufgetreten.
	return 0
}

// GeometricMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das geometrische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das geometrische Mittel ist definiert als
// die n-te Wurzel aus dem Produkt der n Werte.
func GeometricMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	product := float64(intlists.Product(values))
	length := float64(len(values))

	return math.Pow(product, 1.0/length)
}

// HarmonicMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das harmonische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das harmonische Mittel ist definiert als
// die Kehrwert des Durchschnitts der Kehrwerte der n Werte.
func HarmonicMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, v := range values {
		sum += 1.0 / float64(v)
	}

	return float64(len(values)) / sum
}
