package arrays

import (
	"fmt"
)

func DemoArrays1() {
	var numbers [5]int

	fmt.Println(numbers)
}

func DemoArrays2() {
	var numbers [5]int
	numbers[2] = 30
	fmt.Println(numbers)
}

func DemoArrays3() {
	var city [5]string
	city[0] = "Samsun"
	city[1] = "Bursa"
	city[2] = "Manisa"
	city[3] = "Adana"
	city[4] = "İstanbul"
	//sehirler[5] = "Samsun" dizinin tanımlanan max limitini aşıyoruz.

	fmt.Println(city)
	for i := 0; i < 5; i++ {
		fmt.Println(city[i])
	}
}

func DemoArrays4() {
	numbers := [5]int{100, 20, 35, 450, 50}

	//fmt.Println(numbers)

	theBiggest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > theBiggest {
			theBiggest = numbers[i]
		}
		//fmt.Println(sayilar)
	}
	fmt.Println("En Büyül :", theBiggest)
}

func DemoArrays5() {
	var numbersTable [2][3]int

	numbersTable[0][0] = 1
	numbersTable[0][1] = 2
	numbersTable[0][2] = 3
	numbersTable[1][1] = 4
	numbersTable[1][2] = 5

	fmt.Println(numbersTable[1][1])
}

func DemoArrays6() {
	var numbersTable [2][3]int

	numbersTable[0][0] = 1
	numbersTable[0][1] = 3
	numbersTable[0][2] = 5
	numbersTable[1][0] = 2
	numbersTable[1][1] = 4
	numbersTable[1][2] = 6

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(numbersTable[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}

func DemoArrays7() {
	var numbersTable [2][3]string

	numbersTable[0][0] = "Başlık1"
	numbersTable[0][1] = "Başlık2"
	numbersTable[0][2] = "Başlık2"
	numbersTable[1][0] = "Data1"
	numbersTable[1][1] = "Data2"
	numbersTable[1][2] = "Data3"

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(numbersTable[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
