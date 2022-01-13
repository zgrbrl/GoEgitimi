package loops

import "fmt"

func Demo1() {
	metin := "Hello Özgür,"

	i := 1

	for i <= 5 {
		fmt.Printf("%v. %v \n", i, metin)
		i = i + 1
	}

	fmt.Println("Bitti...")
}
