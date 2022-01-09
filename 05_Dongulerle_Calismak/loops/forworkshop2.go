package loops

import "fmt"

func Demo4() {

	aklimdakiSayi := 45
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Lütfen Bir Sayı Giriniz...")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Lütfen Büyük Bir Sayı Giriniz...")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}

		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Lütfen Küçük Bir Sayı Giriniz...")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}

	basariDurumu := ""

	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Println("Bravo Bildiniz...")

	fmt.Printf("%v. deneme sonucunda buldunuz. \nBaşarı Durumunuz: %v", tahminSayisi, basariDurumu)

}
