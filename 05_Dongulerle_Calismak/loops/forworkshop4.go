package loops

import "fmt"

func Demo5() {

	sayi := 0
	fmt.Println("Bir sayı giriniz...")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println("Asaldır.")
	} else {
		fmt.Println("Asal değildir.")
	}

}
