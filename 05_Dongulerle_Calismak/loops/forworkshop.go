package loops

import "fmt"

func Demo3() {
	//Bir değişkende sayı tut.
	//

	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir Sayı Tahmin Ediniz...")
	fmt.Scanln(&tahminEdilenSayi)

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Lütfen büyük bir sayı giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
		}

		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Lütfen küçük bir sayı giriniz.")
			fmt.Scanln(&tahminEdilenSayi)
		}
	}
	if tahminEdilenSayi == aklimdakiSayi {
		fmt.Println("Bravo bildiniz.")
	}

}
