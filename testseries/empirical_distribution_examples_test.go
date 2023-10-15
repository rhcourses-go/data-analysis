package testseries

func ExampleEmpiricalDistribution() {
	absFreqs1 := []int{3, 2, 1, 2, 2, 0, 1}
	relFreqs1 := RelativeFrequencies(absFreqs1)
	empDist1 := EmpiricalDistribution(relFreqs1)
	printFloatList(empDist1)

	absFreqs2 := []int{2, 0, 2, 1, 0, 1}
	relFreqs2 := RelativeFrequencies(absFreqs2)
	empDist2 := EmpiricalDistribution(relFreqs2)
	printFloatList(empDist2)

	// Output:
	// 0.27 0.45 0.55 0.73 0.91 0.91 1.00
	// 0.33 0.33 0.67 0.83 0.83 1.00
}
