package conditionals

import "fmt"

func Demo3() {

	sayi, sayi1, sayi2 := 10, 2, 3

	enbuyuk := sayi

	if sayi1 > enbuyuk {
		enbuyuk = sayi1
	}

	if sayi2 > enbuyuk {
		enbuyuk = sayi2
	}

	fmt.Printf("En büyük sayı: %v\n", enbuyuk)
}
