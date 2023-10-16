package testseries

import "fmt"

// printFloatList gibt eine Liste von Gleitkommazahlen auf der Konsole aus.
// Die Zahlen werden durch Leerzeichen getrennt, am Ende wird ein Zeilenumbruch ausgegeben
// und die Funktion sorgt dafür, dass jede Zahl mit zwei Nachkommastellen ausgegeben wird.
// Dies ist eine Hilfsfunktion für die Tests, damit die Ausgaben nicht zu lang werden.
func printFloatList(list []float64) {
	for _, v := range list[:len(list)-1] {
		fmt.Printf("%.2f ", v)
	}
	fmt.Printf("%.2f\n", list[len(list)-1])
}
