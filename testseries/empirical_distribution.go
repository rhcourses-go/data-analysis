package testseries

// EmpiricalDistribution erwartet eine Liste mit relativen Häufigkeiten einer Messreihe.
// Die Funktion liefert eine Liste, in der für jede Zahl der
// entsprechende Wert der empirischen Verteilungsfunktion steht.
func EmpiricalDistribution(relativeFreqs []float64) []float64 {
	emp := make([]float64, len(relativeFreqs))
	sum := 0.0

	for i, f := range relativeFreqs {
		sum += f
		emp[i] = sum
	}

	return emp
}

// Distribution erwartet eine Liste mit ganzzahligen Messwerten.
// Die Funktion liefert eine Liste mit den Werten der empirischen Verteilungsfunktion.
func Distribution(values []int) []float64 {
	absolute := AbsoluteFrequencies(values)
	relative := RelativeFrequencies(absolute)
	return EmpiricalDistribution(relative)
}
