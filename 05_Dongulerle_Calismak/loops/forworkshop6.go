package loops

import "fmt"

func MultiplicationTable() {

	for factor1 := 1; factor1 <= 10; factor1++ {
		for factor2 := 1; factor2 <= 10; factor2++ {
			fmt.Printf("%d * %d = %d \n", factor1, factor2, (factor1 * factor2))
		}
	}
}
