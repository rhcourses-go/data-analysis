package intlists

// Min berechnet das Minimum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Min(values []int) int {
	min := values[0]
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um das Minimum zu berechnen.
	   Jedes Mal, wenn Sie ein neues Minimum finden, speichern Sie es in der Variablen min.
	*/
	// tag::solution[]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	// end::solution[]
	return min
}

// Max berechnet das Maximum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Max(values []int) int {
	max := values[0]
	/* Hinweis:
	   Gehen Sie analog zu Min vor.
	*/
	// tag::solution[]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	// end::solution[]
	return max
}

// ValueRange erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert eine lückenlose Liste mit allen Zahlen zwischen
// dem Minimum und dem Maximum der Messreihe.
func ValueRange(values []int) []int {
	result := []int{}
	/* Hinweis:
	   Verwenden Sie die Funktionen Min und Max, um das Minimum und das Maximum
	   der Messreihe zu berechnen.
	   Fügen Sie dann in einer Schleife alle Zahlen zwischen
	   Minimum und Maximum zu result hinzu.
	*/
	// tag::solution[]
	min, max := Min(values), Max(values)
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	// end::solution[]
	return result

}

// Sum erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert die Summe aller Werte.
func Sum(values []int) int {
	sum := 0
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Summe zu berechnen.
	   Addieren Sie in jedem Schleifendurchlauf den aktuellen Wert zur Summe.
	*/
	// tag::solution[]
	for _, v := range values {
		sum += v
	}
	// end::solution[]
	return sum
}

// Product erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das Produkt aller Werte.
func Product(values []int) int {
	product := 1
	/* Hinweis:
	   Gehen Sie analog zu Sum vor.
	*/
	// tag::solution[]
	for _, v := range values {
		product *= v
	}
	// end::solution[]
	return product
}
