package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	fmt.Printf("Hesabınızda Bulunan Tutar: %v TL\n", hesap)

	if cekilmekIstenen > hesap {
		fmt.Println("Hesabınızdaki para yetersiz")

	} else if cekilmekIstenen == hesap {
		fmt.Printf("Paranız Hazılanıyor. \nÇekilen Tutar: %v TL\n", hesap)
		fmt.Println("Dikkat, Hesabınızda Para Kalmadı.")
	} else {
		fmt.Println("Paranız Hazırlanıyor.")
	}
}
