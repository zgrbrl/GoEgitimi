package loops

import "fmt"

func HesapMakinasi() {

	operator, toplama, cikarma, carpma, bolme, sayi1, sayi2, sonuc, yeniIslem := "", "+", "-", "*", "/", 0, 0, 0, "E"

	for yeniIslem == "E" {

		fmt.Println("Lütfen operatörlerden bir işlem giriniz. +,-,*,/")
		fmt.Scanln(&operator)

		if operator == toplama {
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi1)
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi2)
			sonuc = sayi1 + sayi2
			fmt.Printf("Toplama İşleminin Sonucu %v\n", sonuc)
		}

		if operator == cikarma {
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi1)
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi2)
			sonuc = sayi1 - sayi2
			fmt.Printf("Toplama İşleminin Sonucu %v\n", sonuc)
		}

		if operator == carpma {
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi1)
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi2)
			sonuc = sayi1 * sayi2
			fmt.Printf("Toplama İşleminin Sonucu %v\n", sonuc)
		}

		if operator == bolme {
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi1)
			fmt.Println("Lütfen işlem için sayı giriniz...")
			fmt.Scanln(&sayi2)
			sonuc = sayi1 / sayi2
			fmt.Printf("Toplama İşleminin Sonucu %v\n", sonuc)
		}
		fmt.Println("Yeni bir işlem yapmak ister misiniz. E/H")
		fmt.Scanln(&yeniIslem)
	}

	fmt.Println("İşleminiz sona ermiştir...")

}
