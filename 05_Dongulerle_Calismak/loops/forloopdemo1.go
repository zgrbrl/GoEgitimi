package loops

import "fmt"

func Demo1() {
	metin := "Hello Özgür,"

	i := 1

	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}

	fmt.Println("Bitti...")
}
