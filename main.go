package main

import "fmt"

func ActivateGiftCard() func(int) int {
	amount := 100
	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}

func main() {
	useGiftCard1 := ActivateGiftCard()
	useGiftCard2 := ActivateGiftCard()
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))
}
