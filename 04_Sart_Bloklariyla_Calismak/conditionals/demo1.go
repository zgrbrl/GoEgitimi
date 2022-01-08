package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesapdaki Para Yetersiz...")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız Hazırlanıyor...")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Printf("Bitti. Hesaptaki Para : %v\n", hesap) //kestirme olan

	fmt.Println("Bitti. Hesaptaki para : " + fmt.Sprintf("%v", hesap)) // uzun yol
}
