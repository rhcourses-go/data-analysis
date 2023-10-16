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
	/* Hinweis:
	   Verwenden Sie die Funktion Sum aus dem Paket intlists.
	*/
	// tag::solution[]
	return float64(intlists.Sum(values)) / float64(len(values))
	// end::solution[]
	// taskreturn: 0.0
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

	/* Hinweis:
	   Verwenden Sie die Funktion sort aus dem Paket slices, um die Liste zu sortieren.
	   Die beiden Anweisungen oben kopieren Sie die Liste vorher, damit die ursprüngliche Liste nicht verändert wird.
	   Bestimmen Sie dann das mittlere Element, oder, falls die Liste eine gerade Anzahl von Elementen hat,
	   den Durchschnitt der beiden mittleren Elemente.
	*/
	// tag::solution[]
	slices.Sort(sorted)
	if len(sorted)%2 == 0 {
		return (sorted[len(sorted)/2-1] + sorted[len(sorted)/2]) / 2
	}
	// end::solution[]
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
	maxpos := 0
	/* Hinweis:
	   Bestimmen Sie den kleinsten Wert, der in der Liste vorkommt.
	   Bestimmen Sie außerdem die Liste der absoluten Häufigkeiten und deren Maximum.
	   Dies ist der Wert, der am häufigsten vorkommt.
	   Suchen Sie dann nach dessen Position in der Liste der absoluten Häufigkeiten.
	*/

	// tag::solution[]
	// Bestimme den kleinsten Wert in der Liste der Werte.
	minvalue := intlists.Min(values)

	// Bestimme, wie oft der häufigste Wert vorkommt.
	// Bestimme dazu die Liste der absoluten Häufigkeiten und deren Maximum.
	freq := AbsoluteFrequencies(values)
	max := intlists.Max(freq)

	// Suche die Position dieses Wertes in der Liste der absoluten Häufigkeiten.
	// Aus dieser Position kann der häufigste Wert bestimmt werden.
	for i, f := range freq {
		if f == max {
			maxpos = i + minvalue
			break
		}
	}
	// end::solution[]
	return maxpos
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
	/* Hinweis:
	   Verwenden Sie die Funktion Product aus dem Paket intlists.
	   Verwenden Sie die Funktion Pow aus dem Paket math.
	*/
	// tag::solution[]
	product := float64(intlists.Product(values))
	length := float64(len(values))

	return math.Pow(product, 1.0/length)
	// end::solution[]
	// taskreturn: 0.0
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

	/* Hinweis:
	   Bestimmen Sie die Summe der Kehrwerte der Werte.
	   Teilen Sie die Anzahl der Werte durch diese Summe.
	*/
	sum := 0.0
	for _, v := range values {
		sum += 1.0 / float64(v)
	}

	return float64(len(values)) / sum
}
