package main

import "fmt"

func main() {

	var metin string = "Hello Özgür"
	//metin := "Hello Özgür"
	fmt.Println(metin)

	var kdv int = 10
	//kdv := 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * kdv / 100))

	var kdv2 float32 = 0.1
	//kdv2 := 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25
	fmt.Println(kdv3)
	//Veri Tipini Öğrenme
	fmt.Printf("Veri Tipi: %T\n", kdv3)

	var durum bool

	var firstName string = "Özgür"
	var lastName string = "Birel"

	//şart blokları "!=" durum eşit değil ise false
	//şart blokları "==" durum eşit ise true

	durum = firstName != lastName

	fmt.Println(durum)
	fmt.Println(!durum)

}
